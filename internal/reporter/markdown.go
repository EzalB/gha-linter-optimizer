package reporter

import (
	"fmt"
	"strings"

	"github.com/EzalB/gha-linter-optimizer/internal/rules"
)

func GenerateMarkdownReport(issues []rules.Issue) string {
	if len(issues) == 0 {
		return "✅ No lint issues found!"
	}

	var sb strings.Builder
	sb.WriteString("### 🚨 GHA Lint Report\n\n")
	// for _, r := range results {
	// 	sb.WriteString(fmt.Sprintf("- %s\n", r))
	// }

	for _, i := range issues {
		sb.WriteString(
			fmt.Sprintf(
				"- **%s**: %s _(📄 %s:%d)_\n",
				i.Rule, i.Message, i.File, i.Line,
			),
		)
	}

	return sb.String()
}
