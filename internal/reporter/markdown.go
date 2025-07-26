package reporter

import (
	"fmt"
	"strings"
)

func GenerateMarkdownReport(results []string) string {
	if len(results) == 0 {
		return "✅ No lint issues found!"
	}

	var sb strings.Builder
	sb.WriteString("### 🚨 GHA Lint Report\n\n")
	for _, r := range results {
		sb.WriteString(fmt.Sprintf("- %s\n", r))
	}
	return sb.String()
}
