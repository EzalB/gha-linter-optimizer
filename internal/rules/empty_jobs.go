package rules

import (
	"fmt"

	"github.com/EzalB/gha-linter-optimizer/internal/parser"
)

type EmptyJobsRule struct{}

func (r EmptyJobsRule) Name() string {
	return "Empty Jobs Rule"
}
func (r EmptyJobsRule) Description() string {
	return "Detect jobs with no steps"
}

func (r EmptyJobsRule) Apply(wf *parser.Workflow) []Issue {
	var issues []Issue
	
	for jobID, job := range wf.Jobs {
		if len(job.Steps) == 0 {
			issues = append(issues, Issue{
				Rule:     r.Name(),
				Message:  fmt.Sprintf("Job '%s' has no steps defined", jobID),
				File:     wf.FilePath,  // from parser.Workflow
				Line:     job.Line,     // from parser.Job
				Severity: "warning",
			})
		}
	}
	return issues
}
