package rules

import (
	"fmt"
	"regexp"

	"github.com/EzalB/gha-linter-optimizer/internal/parser"
)

type LatestTagRule struct{}

func (r LatestTagRule) Name() string       { return "Latest Tag Rule" }
func (r LatestTagRule) Description() string { return "Disallow @latest in reusable actions" }

func (r LatestTagRule) Apply(wf *parser.Workflow) []string {
	var warnings []string
	reg := regexp.MustCompile(`@latest$`)
	for jobID, job := range wf.Jobs {
		for _, step := range job.Steps {
			if reg.MatchString(step.Uses) {
				warnings = append(warnings, fmt.Sprintf("Job '%s': step uses '%s' (avoid @latest)", jobID, step.Uses))
			}
		}
	}
	return warnings
}
