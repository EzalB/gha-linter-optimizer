package rules

import (
	"fmt"

	"github.com/EzalB/gha-linter-optimizer/internal/parser"
)

type MissingRequiredFieldsRule struct{}

func (r MissingRequiredFieldsRule) Name() string {
	return "Missing Required Fields Rule"
}
func (r MissingRequiredFieldsRule) Description() string {
	return "Ensure required fields like runs-on and steps exist"
}

func (r MissingRequiredFieldsRule) Apply(wf *parser.Workflow) []Issue  {
	var issues []Issue

	for jobID, job := range wf.Jobs {
		if job.RunsOn == "" {
			issues = append(issues, Issue{
				Rule:     r.Name(),
				Message:  fmt.Sprintf("Job '%s' is missing the required 'runs-on' field", jobID),
				File:     wf.FilePath,
				Line:     job.Line,
				Severity: "error",
			})
		}

		if len(job.Steps) == 0 {
			issues = append(issues, Issue{
				Rule:     r.Name(),
				Message:  fmt.Sprintf("Job '%s' is missing the required 'steps' section", jobID),
				File:     wf.FilePath,
				Line:     job.Line,
				Severity: "error",
			})
		}
	}
	return issues
}
