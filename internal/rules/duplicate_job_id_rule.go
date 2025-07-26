package rules

import (
	"fmt"
	"github.com/EzalB/gha-linter-optimizer/internal/parser"
)

type DuplicateJobIDRule struct{}

func (r DuplicateJobIDRule) Name() string       { return "Duplicate Job ID Rule" }
func (r DuplicateJobIDRule) Description() string { return "Detect duplicate job IDs" }

func (r DuplicateJobIDRule) Apply(wf *parser.Workflow) []string {
	var warnings []string
	seen := make(map[string]bool)
	for jobID := range wf.Jobs {
		if seen[jobID] {
			warnings = append(warnings, fmt.Sprintf("Duplicate job ID found: '%s'", jobID))
		}
		seen[jobID] = true
	}
	return warnings
}
