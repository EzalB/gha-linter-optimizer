package parser

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Workflow struct {
	Jobs map[string]Job `yaml:"jobs"`
}

type Job struct {
	Name  	string 				`yaml:"name"`
	RunsOn  string 				`yaml:"runs"`
	Steps 	[]Step 				`yaml:"steps"`
	Needs   []string 			`yaml:"needs,omitempty"`
	Env     map[string]string 	`yaml:"env,omitempty"`
}

type Step struct {
	Name string `yaml:"name"`
	Uses string `yaml:"uses"`
	Run  string `yaml:"run-on"`
	If   string `yaml:"if"`
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
