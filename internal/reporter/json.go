package reporter

import (
	"encoding/json"
)

func GenerateJSONReport(results []string) (string, error) {
	data := map[string]interface{}{
		"issues": results,
		"count":  len(results),
	}
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
