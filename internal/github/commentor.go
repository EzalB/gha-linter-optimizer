package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/EzalB/gha-linter-optimizer/internal/rules"
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
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		panic("GITHUB_TOKEN not set")
	}
	
	return &GitHubCommenter{
		Token:     token,
		Repo:      repo,
		PRNumber:  prNumber,
		APIURL:    "https://api.github.com",
		UserAgent: "gha-linter-bot",
	}
}

func (g *GitHubCommenter) BuildPRComment(issues []rules.Issue) string {
	if len(issues) == 0 {
		return botMarker + "\n" + "✅ No workflow issues found!"
	}

	var b strings.Builder
	b.WriteString(botMarker + "\n")
	b.WriteString("### 🚨 GitHub Actions Lint Report\n\n")

	for _, i := range issues {
		b.WriteString(fmt.Sprintf(
			"- **%s**: %s _(📄 `%s`:%d)_\n",
			i.Rule, i.Message, i.File, i.Line,
		))
	}

	return b.String()
}

func (g *GitHubCommenter) PostOrUpdateComment(commentBody string) error {
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

	//commentBody := botMarker + "\n" + markdownReport
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
