package rules

import "github.com/EzalB/gha-linter-optimizer/internal/parser"

type Rule interface {
	Name() string
	Description() string
	Apply(wf *parser.Workflow) []string
}
