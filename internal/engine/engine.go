package engine

import (
	"path/filepath"

	"github.com/EzalB/gha-linter-optimizer/internal/parser"
	"github.com/EzalB/gha-linter-optimizer/internal/rules"
	"github.com/EzalB/gha-linter-optimizer/internal/utils"
)

func RunLint(path string) []string {
	files := utils.GetWorkflowFiles(path)
	allResults := []string{}

	for _, file := range files {
		wf, err := parser.ParseWorkflow(file)
		if err != nil {
			continue
		}

		activeRules := []rules.Rule{
			rules.LatestTagRule{},
			rules.CacheRule{},
			rules.NameRule{},
		}

		for _, rule := range activeRules {
			results := rule.Apply(wf)
			for _, res := range results {
				allResults = append(allResults, filepath.Base(file)+": "+res)
			}
		}
	}
	return allResults
}
