package engine

import (
	"path/filepath"
	"fmt"

	"github.com/EzalB/gha-linter-optimizer/internal/parser"
	"github.com/EzalB/gha-linter-optimizer/internal/rules"
	"github.com/EzalB/gha-linter-optimizer/internal/utils"
)

func GetAllRules() []rules.Rule {
	return []rules.Rule{
		// rules.LatestTagRule{},
		rules.NameRule{},
		// rules.EmptyJobsRule{},
		// rules.DuplicateJobIDRule{},
		// rules.ZombieStepsRule{},
		// rules.BrokenReferencesRule{},
		// rules.MissingRequiredFieldsRule{},
		// rules.SecretsExposureRule{},
	}
}

func RunLint(path string) []rules.Issue {
	files := parser.GetWorkflowFiles(path)
	var all []rules.Issue
	// allResults := []string{}

	for _, file := range files {
		wf, err := parser.ParseWorkflow(file)
		if err != nil {
			utils.Log.Warn("❌ Failed to parse workflow", "file", file, "error", err.Error())
			// all = append(all, fmt.Sprintf("Parsing error in %s: %v", filepath.Base(file), err))
			continue
		}

		for _, rule := range GetAllRules() {
			issues := rule.Apply(wf)
			all = append(all, issues...)
		}
	}
	return all
}
