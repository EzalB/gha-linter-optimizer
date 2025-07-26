package reporter

import (
	"encoding/json"
)

func GenerateJSONReport(issues []string) (string, error) {
	out, err := json.MarshalIndent(map[string]interface{}{
		"issues": issues,
	}, "", "  ")
	return string(out), err
}
