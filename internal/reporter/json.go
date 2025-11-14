package reporter

import (
	"encoding/json"

	"github.com/EzalB/gha-linter-optimizer/internal/rules"
)

func GenerateJSONReport(issues []rules.Issue) (string, error) {
	data := map[string]interface{}{
		"issues": issues,
		"count":  len(issues),
	}
	
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
