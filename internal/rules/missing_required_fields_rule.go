package rules

import (
	"fmt"

	"github.com/EzalB/gha-linter-optimizer/internal/parser"
)

type MissingRequiredFieldsRule struct{}

func (r MissingRequiredFieldsRule) Name() string       { return "Missing Required Fields Rule" }
func (r MissingRequiredFieldsRule) Description() string { return "Ensure required fields like runs-on and steps exist" }

func (r MissingRequiredFieldsRule) Apply(wf *parser.Workflow) []string {
	var warnings []string
	for jobID, job := range wf.Jobs {
		if job.RunsOn == "" {
			warnings = append(warnings, fmt.Sprintf("Job '%s' missing 'runs-on' field", jobID))
		}
		if len(job.Steps) == 0 {
			warnings = append(warnings, fmt.Sprintf("Job '%s' missing 'steps' section", jobID))
		}
	}
	return warnings
}
