package rules

import (
	"fmt"

	"github.com/EzalB/gha-linter-optimizer/internal/parser"
)

type NameRule struct{}

func (r NameRule) Name() string       { return "Name Rule" }
func (r NameRule) Description() string { return "Ensure job and step names are present" }

func (r NameRule) Apply(wf *parser.Workflow) []string {
	var results []string
	for jobID, job := range wf.Jobs {
		if job.Name == "" {
			results = append(results, fmt.Sprintf("Job '%s' has no name", jobID))
		}
		for _, step := range job.Steps {
			if step.Name == "" && step.Uses == "" {
				results = append(results, fmt.Sprintf("Job '%s' has unnamed step", jobID))
			}
		}
	}
	return results
}
