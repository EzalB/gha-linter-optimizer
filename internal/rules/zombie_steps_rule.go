package rules

import (
	"fmt"
	"github.com/EzalB/gha-linter-optimizer/internal/parser"
)

type ZombieStepsRule struct{}

func (r ZombieStepsRule) Name() string {
	return "Zombie Steps Rule"
}
func (r ZombieStepsRule) Description() string {
	return "Detect steps that are defined but never triggered"
}

func (r ZombieStepsRule) Apply(wf *parser.Workflow) []Issue {
	var issues []Issue
	
	for jobID, job := range wf.Jobs {
		for _, step := range job.Steps {
			isAlwaysDisabled := step.If == "false"
			isEmpty := step.Name == "" && step.Run == "" && step.Uses == ""

			if isAlwaysDisabled || isEmpty {
				issues = append(issues, Issue{
					Rule:     r.Name(),
					Message:  fmt.Sprintf("Job '%s': step '%s' appears to be a zombie step (never triggered or empty)", jobID, step.Name),
					File:     wf.FilePath,
					Line:     step.Line,
					Severity: "warning",
				})
			}
		}
	}
	return issues
}
