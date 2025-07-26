package rules

import (
	"fmt"
	"strings"

	"github.com/EzalB/gha-linter-optimizer/internal/parser"
)

type CacheRule struct{}

func (r CacheRule) Name() string       { return "Cache Rule" }
func (r CacheRule) Description() string { return "Recommend usage of actions/cache for dependencies" }

func (r CacheRule) Apply(wf *parser.Workflow) []string {
	var findings []string
	for jobID, job := range wf.Jobs {
		usesCache := false
		for _, step := range job.Steps {
			if strings.Contains(step.Uses, "actions/cache") {
				usesCache = true
			}
		}
		if !usesCache {
			findings = append(findings, fmt.Sprintf("Job '%s' does not use 'actions/cache'", jobID))
		}
	}
	return findings
}
