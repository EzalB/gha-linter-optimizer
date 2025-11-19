package rules

import (
	"fmt"
	"regexp"

	"github.com/EzalB/gha-linter-optimizer/internal/parser"
)

type LatestTagRule struct{}

func (r LatestTagRule) Name() string {
	return "Latest Tag Rule"
}
func (r LatestTagRule) Description() string {
	return "Detects @latest tag for steps in actions"
}

func (r LatestTagRule) Apply(wf *parser.Workflow) []Issue {
	var issues []Issue
	
	reg := regexp.MustCompile(`@latest$`)
	
	for jobID, job := range wf.Jobs {
		for _, step := range job.Steps {
			if reg.MatchString(step.Uses) {
				issues = append(issues, Issue{
					Rule:     r.Name(),
					Message:  fmt.Sprintf("Job '%s': step uses '%s' (avoid @latest)", jobID, step.Uses),
					File:     wf.FilePath,
					Line:     step.Line,
					Severity: "warning",
				})
			}
		}
	}
	return issues
}
