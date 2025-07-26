package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/EzalB/gha-linter-optimizer/internal/config"
	"github.com/EzalB/gha-linter-optimizer/internal/engine"
	"github.com/EzalB/gha-linter-optimizer/internal/github"
	"github.com/EzalB/gha-linter-optimizer/internal/reporter"
	"github.com/EzalB/gha-linter-optimizer/internal/utils"
)

func main() {
	utils.InitLogger()
	utils.Log.Info("🚀 Starting GitHub Actions Linter & Optimizer...")

	// Load config
	cfg := config.Load()
	if cfg.Path == "" {
		utils.Log.Error("❌ Workflow path not specified.")
		os.Exit(1)
	}

	// Run linting engine
	results := engine.RunLint(cfg.Path)

	// Generate report
	markdownReport := reporter.GenerateMarkdownReport(results)
	jsonReport, err := reporter.GenerateJSONReport(results)
	if err != nil {
		utils.Log.Errorf("Failed to generate JSON report: %v", err)
	}


	// Print report to stdout
	if len(results) > 0 {
		fmt.Println("## 🔍 GitHub Actions Lint Report (Markdown)")
		fmt.Println(markdownReport)
		fmt.Println("## 🔍 GitHub Actions Lint Report (JSON)")
		fmt.Println(string(jsonReport))
	} else {
		fmt.Println("✅ No issues found in GitHub Actions workflows!")
	}

	// PR comment integration (if env vars are set)
	repo := os.Getenv("GITHUB_REPOSITORY")
	prNumStr := os.Getenv("PR_NUMBER") // Usually passed in workflow_dispatch or issue_comment
	githubToken := os.Getenv("GITHUB_TOKEN")

	if repo != "" && prNumStr != "" && githubToken != "" {
		prNum, err := strconv.Atoi(prNumStr)
		if err != nil {
			utils.Log.Errorf("Invalid PR number: %s", prNumStr)
			os.Exit(1)
		}

		commenter := github.NewGitHubCommenter(repo, strconv.Itoa(prNum))
		err = commenter.PostOrUpdateComment(markdownReport)
		if err != nil {
			utils.Log.Errorf("Failed to post comment: %v", err)
			os.Exit(1)
		}
		utils.Log.Info("💬 PR comment posted successfully.")
	} else {
		utils.Log.Info("Skipping PR comment (missing GITHUB_REPOSITORY or PR_NUMBER or GITHUB_TOKEN)")
	}

	// Exit non-zero if issues found
	if len(results) > 0 {
		os.Exit(1)
	}
}
