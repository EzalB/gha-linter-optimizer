package parser

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Workflow struct {
	Jobs map[string]Job `yaml:"jobs"`
}

type Job struct {
	Name  string `yaml:"name"`
	Steps []Step `yaml:"steps"`
}

type Step struct {
	Name string `yaml:"name"`
	Uses string `yaml:"uses"`
	Run  string `yaml:"run"`
}

func ParseWorkflow(path string) (*Workflow, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var wf Workflow
	if err := yaml.Unmarshal(data, &wf); err != nil {
		return nil, err
	}
	return &wf, nil
}
