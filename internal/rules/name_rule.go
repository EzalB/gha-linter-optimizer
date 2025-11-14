package rules

import (
	"fmt"

	"github.com/EzalB/gha-linter-optimizer/internal/parser"
)

type NameRule struct{}

func (r NameRule) Name() string       { return "Name Rule" }
func (r NameRule) Description() string { return "Ensure job and step names are present" }

func (r NameRule) Apply(wf *parser.Workflow) []Issue {
	var results []Issue
	for jobID, job := range wf.Jobs {
		if job.Name == "" {
			results = append(results, Issue{
				Rule:    r.Name(),
				File:    wf.Path,
				Line:    job.Line,
				Message: fmt.Sprintf("Job '%s' has no name", jobID),
				Severity: "warning",
			})
		}
		for _, step := range job.Steps {
			if step.Name == "" && step.Uses == "" {
				results = append(results, Issue{
					Rule:    r.Name(),
					File:    wf.Path,
					Line:    step.Line,
					Message: fmt.Sprintf("Job '%s': step '%s' has no name", jobID, step.Run),
					Severity: "warning",
				})
			}
		}
	}
	return results
}
