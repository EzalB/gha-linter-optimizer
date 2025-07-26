package rules

import (
	"fmt"
	"github.com/EzalB/gha-linter-optimizer/internal/parser"
)

type ZombieStepsRule struct{}

func (r ZombieStepsRule) Name() string       { return "Zombie Steps Rule" }
func (r ZombieStepsRule) Description() string { return "Detect steps that are defined but never triggered" }

func (r ZombieStepsRule) Apply(wf *parser.Workflow) []string {
	var warnings []string
	for jobID, job := range wf.Jobs {
		for _, step := range job.Steps {
			if step.If == "false" || step.Name == "" && step.Run == "" && step.Uses == "" {
				warnings = append(warnings, fmt.Sprintf("Job '%s': possibly zombie step with no run/uses/if condition always false", jobID))
			}
		}
	}
	return warnings
}
