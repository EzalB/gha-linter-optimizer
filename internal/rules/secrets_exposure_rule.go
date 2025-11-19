package rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/EzalB/gha-linter-optimizer/internal/parser"
)

type SecretsExposureRule struct{}

func (r SecretsExposureRule) Name() string       { return "Secrets Exposure Rule" }
func (r SecretsExposureRule) Description() string { return "Detects potential secret, password, or token exposures in workflows" }

func (r SecretsExposureRule) Apply(wf *parser.Workflow) []Issue {
	var issues []Issue

	// Common secret patterns (tokens, keys, passwords, API keys, etc.)
	patterns := []*regexp.Regexp{
		regexp.MustCompile(`(?i)(password|secret|token|apikey|api_key|accesskey|aws_access_key_id|aws_secret_access_key|github_token)\s*[:=]\s*['"]?[A-Za-z0-9_\-]{8,}['"]?`),
		regexp.MustCompile(`ghp_[A-Za-z0-9]{30,}`),                              // GitHub PAT
		regexp.MustCompile(`gho_[A-Za-z0-9]{30,}`),                              // GitHub OAuth token
		regexp.MustCompile(`ghs_[A-Za-z0-9]{30,}`),                              // GitHub App installation token
		regexp.MustCompile(`ghr_[A-Za-z0-9]{30,}`),                              // GitHub Refresh token
		regexp.MustCompile(`AKIA[0-9A-Z]{16}`),                                  // AWS Access Key
		regexp.MustCompile(`(?i)authorization:\s*Bearer\s+[A-Za-z0-9\.\-_]+`),  // Bearer tokens
		regexp.MustCompile(`(?i)(slack|discord).*hook.*url.*[:=].*https?://[^\s]+`), // Webhook URLs
	}

	for jobID, job := range wf.Jobs {
		for _, step := range job.Steps {
			content := strings.Join([]string{step.Run, step.Uses, step.Raw}, "\n")
			// line := strings.TrimSpace(step.Run + step.Uses + step.Name)
			
			for _, pattern := range patterns {
				if pattern.MatchString(content) {
					issues = append(issues, Issue{
						Rule:     r.Name(),
						Message:  fmt.Sprintf("Job '%s': possible secret or token detected in step '%s'", jobID, step.Name),
						File:     wf.Path,
						Line:     step.Line,
						Severity: "error",
					})
					break
				}
			}

			// Check environment variables defined in steps (possible plaintext secrets)
			for key, val := range step.Env {

				keyLower := strings.ToLower(key)
				if strings.Contains(keyLower, "secret") ||
					strings.Contains(keyLower, "token") ||
					strings.Contains(keyLower, "password") ||
					strings.Contains(keyLower, "key") {

					issues = append(issues, Issue{
						Rule:     r.Name(),
						Message:  fmt.Sprintf("Job '%s': environment variable '%s' may contain sensitive data", jobID, key),
						File:     wf.Path,
						Line:     step.Line, // best available context
						Severity: "error",
					})
				}

				// Also detect if value itself looks like a secret
				for _, pattern := range patterns {
					if pattern.MatchString(val) {
						issues = append(issues, Issue{
							Rule:     r.Name(),
							Message:  fmt.Sprintf("Job '%s': environment variable '%s' contains a possible secret", jobID, key),
							File:     wf.Path,
							Line:     step.Line,
							Severity: "error",
						})
						break
					}
				}
			}
		}
	}

	return issues
}