package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const botMarker = "<!-- gha-linter-bot-comment -->"

type GitHubCommenter struct {
	Token     string
	Repo      string
	PRNumber  string
	APIURL    string
	UserAgent string
}

func NewGitHubCommenter(repo, prNumber string) *GitHubCommenter {
	if os.Getenv("GITHUB_TOKEN") == "" {
		panic("GITHUB_TOKEN environment variable not set")
	}
	
	return &GitHubCommenter{
		Token:     os.Getenv("GITHUB_TOKEN"),
		Repo:      repo,
		PRNumber:  prNumber,
		APIURL:    "https://api.github.com",
		UserAgent: "gha-linter-bot",
	}
}

func (g *GitHubCommenter) PostOrUpdateComment(markdownReport string) error {
	url := fmt.Sprintf("%s/repos/%s/issues/%s/comments", g.APIURL, g.Repo, g.PRNumber)

	// Fetch existing comments
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+g.Token)
	req.Header.Set("User-Agent", g.UserAgent)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	var comments []struct {
		ID   int    `json:"id"`
		Body string `json:"body"`
	}
	json.Unmarshal(bodyBytes, &comments)

	botCommentID := -1
	for _, c := range comments {
		if strings.Contains(c.Body, botMarker) {
			botCommentID = c.ID
			break
		}
	}

	commentBody := botMarker + "\n" + markdownReport
	payload := map[string]string{"body": commentBody}
	payloadBytes, _ := json.Marshal(payload)

	if botCommentID != -1 {
		// Update existing comment
		url = fmt.Sprintf("%s/repos/%s/issues/comments/%d", g.APIURL, g.Repo, botCommentID)
		req, _ = http.NewRequest("PATCH", url, bytes.NewReader(payloadBytes))
	} else {
		// Create new comment
		req, _ = http.NewRequest("POST", url, bytes.NewReader(payloadBytes))
	}

	req.Header.Set("Authorization", "Bearer "+g.Token)
	req.Header.Set("User-Agent", g.UserAgent)
	req.Header.Set("Content-Type", "application/json")
	_, err = http.DefaultClient.Do(req)
	return err
}
