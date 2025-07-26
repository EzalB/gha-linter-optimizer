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
		rules.LatestTagRule{},
		rules.NameRule{},
		rules.EmptyJobsRule{},
		rules.DuplicateJobIDRule{},
		rules.ZombieStepsRule{},
		rules.BrokenReferencesRule{},
		rules.MissingRequiredFieldsRule{},
	}
}

func RunLint(path string) []string {
	files := utils.GetWorkflowFiles(path)
	allResults := []string{}

	for _, file := range files {
		wf, err := parser.ParseWorkflow(file)
		if err != nil {
			utils.Log.Warn("❌ Failed to parse workflow", "file", file, "error", err.Error())
			allResults = append(allResults, fmt.Sprintf("Parsing error in %s: %v", filepath.Base(file), err))
			continue
		}

		for _, rule := range GetAllRules() {
			results := rule.Apply(wf)
			for _, res := range results {
				allResults = append(allResults, filepath.Base(file)+": "+res)
			}
		}
	}
	return allResults
}
