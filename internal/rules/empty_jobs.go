package rules

import (
	"fmt"

	"github.com/EzalB/gha-linter-optimizer/internal/parser"
)

type EmptyJobsRule struct{}

func (r EmptyJobsRule) Name() string       { return "Empty Jobs Rule" }
func (r EmptyJobsRule) Description() string { return "Detect jobs with no steps" }

func (r EmptyJobsRule) Apply(wf *parser.Workflow) []string {
	var warnings []string
	for jobID, job := range wf.Jobs {
		if len(job.Steps) == 0 {
			warnings = append(warnings, fmt.Sprintf("Job '%s' has no steps defined", jobID))
		}
	}
	return warnings
}
