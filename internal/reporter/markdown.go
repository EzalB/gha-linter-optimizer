package reporter

import (
	"fmt"
	"strings"
)

func GenerateMarkdownReport(issues []string) string {
	if len(issues) == 0 {
		return "✅ No lint issues found!"
	}

	var sb strings.Builder
	sb.WriteString("### 🚨 GHA Lint Report\n\n")
	for _, issue := range issues {
		sb.WriteString(fmt.Sprintf("- %s\n", issue))
	}
	return sb.String()
}
