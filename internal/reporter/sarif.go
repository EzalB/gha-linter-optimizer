package reporter

import (
	"encoding/json"

	"github.com/EzalB/gha-linter-optimizer/internal/rules"
)

type SarifLog struct {
	Version string     `json:"version"`
	Runs    []SarifRun `json:"runs"`
}

type SarifRun struct {
	Tool    SarifTool      `json:"tool"`
	Results []SarifResult  `json:"results"`
}

type SarifTool struct {
	Driver SarifDriver `json:"driver"`
}

type SarifDriver struct {
	Name           string   `json:"name"`
	InformationURI string   `json:"informationUri"`
}

type SarifResult struct {
	RuleID   string         `json:"ruleId"`
	Level    string         `json:"level"`
	Message  SarifMessage   `json:"message"`
	Locations []SarifLocation `json:"locations"`
}

type SarifMessage struct {
	Text string `json:"text"`
}

type SarifLocation struct {
	PhysicalLocation SarifPhysicalLocation `json:"physicalLocation"`
}

type SarifPhysicalLocation struct {
	ArtifactLocation SarifArtifactLocation `json:"artifactLocation"`
	Region           SarifRegion           `json:"region"`
}

type SarifArtifactLocation struct {
	URI string `json:"uri"`
}

type SarifRegion struct {
	StartLine int `json:"startLine"`
}

func GenerateSarif(issues []rules.Issue) (string, error) {
	results := []SarifResult{}

	for _, i := range issues {
		level := "warning"
		if i.Severity == "error" {
			level = "error"
		}

		results = append(results, SarifResult{
			RuleID: i.Rule,
			Level:  level,
			Message: SarifMessage{
				Text: i.Message,
			},
			Locations: []SarifLocation{
				{
					PhysicalLocation: SarifPhysicalLocation{
						ArtifactLocation: SarifArtifactLocation{
							URI: i.File,
						},
						Region: SarifRegion{
							StartLine: i.Line,
						},
					},
				},
			},
		})
	}

	sarif := SarifLog{
		Version: "2.1.0",
		Runs: []SarifRun{
			{
				Tool: SarifTool{
					Driver: SarifDriver{
						Name:           "gha-linter",
						InformationURI: "https://github.com/EzalB/gha-linter-optimizer",
					},
				},
				Results: results,
			},
		},
	}

	out, err := json.MarshalIndent(sarif, "", "  ")
	if err != nil {
		return "", err
	}

	return string(out), nil
}
