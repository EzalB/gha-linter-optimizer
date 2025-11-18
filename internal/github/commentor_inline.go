package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/EzalB/gha-linter-optimizer/internal/rules"
)

type InlineCommenter struct {
	Token     string
	Repo      string
	PRNumber  string
	CommitSHA string
	APIURL    string
	UserAgent string
}

func NewInlineCommenter(repo, pr, commitSHA string) *InlineCommenter {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		panic("GITHUB_TOKEN not set")
	}

	return &InlineCommenter{
		Token:     token,
		Repo:      repo,
		PRNumber:  pr,
		CommitSHA: commitSHA,
		APIURL:    "https://api.github.com",
		UserAgent: "gha-linter-inline-bot",
	}
}

type InlineCommentPayload struct {
	Body      string `json:"body"`
	CommitID  string `json:"commit_id"`
	Path      string `json:"path"`
	Line      int    `json:"line"`
	Side      string `json:"side"` // LEFT or RIGHT
}

func (c *InlineCommenter) PostInlineComments(issues []rules.Issue) error {
	for _, i := range issues {
		comment := InlineCommentPayload{
			Body: fmt.Sprintf("**%s**: %s", i.Rule, i.Message),
			CommitID: c.CommitSHA,
			Path: strings.TrimPrefix(i.File, "./"),
			Line: i.Line,
			Side: "RIGHT",
		}

		url := fmt.Sprintf("%s/repos/%s/pulls/%s/comments", 
			c.APIURL, c.Repo, c.PRNumber)

		b, _ := json.Marshal(comment)

		req, _ := http.NewRequest("POST", url, bytes.NewReader(b))
		req.Header.Set("Authorization", "Bearer "+c.Token)
		req.Header.Set("User-Agent", c.UserAgent)
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		resp.Body.Close()
	}

	return nil
}
