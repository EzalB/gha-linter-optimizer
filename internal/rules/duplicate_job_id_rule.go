package rules

import (
	"fmt"
	"github.com/EzalB/gha-linter-optimizer/internal/parser"
)

type DuplicateJobIDRule struct{}

func (r DuplicateJobIDRule) Name() string {
	return "Duplicate Job ID Rule"
}
func (r DuplicateJobIDRule) Description() string {
	return "Detect duplicate job IDs"
}

func (r DuplicateJobIDRule) Apply(wf *parser.Workflow) []Issue {
	var issues []Issue

	// seen := make(map[string]bool)
	seen := make(map[string]int)
	
	for jobID, job := range wf.Jobs {
		if firstLine, exists := seen[jobID]; exists {
			issues = append(issues, Issue{
				Rule:     r.Name(),
				Message:  fmt.Sprintf("Duplicate job ID found: '%s' (first declared at line %d)", jobID, firstLine),
				File:     wf.FilePath,
				Line:     job.Line, // current duplicate line
				Severity: "error",  // duplicate IDs break workflow execution
			})
		} else {
			seen[jobID] = job.Line
		}
	}
	return issues
}
