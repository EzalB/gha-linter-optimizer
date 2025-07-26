package rules

import (
	"fmt"
	"strings"

	"github.com/EzalB/gha-linter-optimizer/internal/parser"
)

type BrokenReferencesRule struct{}

func (r BrokenReferencesRule) Name() string       { return "Broken References Rule" }
func (r BrokenReferencesRule) Description() string { return "Detect malformed or suspicious run/uses references" }

func (r BrokenReferencesRule) Apply(wf *parser.Workflow) []string {
	var warnings []string
	for jobID, job := range wf.Jobs {
		for _, step := range job.Steps {
			if step.Uses != "" && !strings.Contains(step.Uses, "/") {
				warnings = append(warnings, fmt.Sprintf("Job '%s': suspicious uses: '%s' (missing org/repo?)", jobID, step.Uses))
			}
		}
	}
	return warnings
}
