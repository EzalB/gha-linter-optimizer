package rules

import (
	"fmt"
	"strings"

	"github.com/EzalB/gha-linter-optimizer/internal/parser"
)

type CacheRule struct{}

func (r CacheRule) Name() string {
	return "Cache Rule"
}
func (r CacheRule) Description() string {
	return "Recommend usage of actions/cache for dependencies"
}

func (r CacheRule) Apply(wf *parser.Workflow) []Issue {
	var issues []Issue

	for jobID, job := range wf.Jobs {
		usesCache := false
		
		for _, step := range job.Steps {
			if strings.Contains(step.Uses, "actions/cache") {
				usesCache = true
				break
			}
		}
		
		if !usesCache {
			issues = append(issues, Issue{
				Rule:     r.Name(),
				Message:  fmt.Sprintf("Job '%s' does not use 'actions/cache' (recommended to speed up CI)", jobID),
				File:     wf.FilePath,
				Line:     job.Line, // Best available line number for the job
				Severity: "info",
			})
		}
	}

	return issues
}
