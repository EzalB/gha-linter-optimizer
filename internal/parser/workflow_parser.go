package parser

import (
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Workflow struct {
	Jobs 		map[string]Job `yaml:"jobs"`
	RawLines 	[]string		`yaml:"-"`
	Path 		string			`yaml:"-"`
}

type Job struct {
	Name  	string 				`yaml:"name"`
	Uses	string				`yaml:"uses"`
	RunsOn 	string   			`yaml:"runs-on"`
	Steps 	[]Step 				`yaml:"steps"`
	Needs   []string 			`yaml:"needs,omitempty"`
	Env     map[string]string 	`yaml:"env,omitempty"`
}

type Step struct {
	Name 	string 				`yaml:"name"`
	Uses 	string 				`yaml:"uses"`
	Run  	string 				`yaml:"run"`
	If   	string 				`yaml:"if"`
	Env  	map[string]string 	`yaml:"env,omitempty"`
}

func ParseWorkflow(path string) (*Workflow, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var wf Workflow
	wf.Path = path
	wf.RawLines = strings.Split(string(data), "\n")
	if err := yaml.Unmarshal(data, &wf); err != nil {
		return nil, err
	}
	return &wf, nil
}

func (wf *Workflow) FindLineNumber(search string) int {
	for i, line := range wf.RawLines {
		if strings.Contains(line, search) {
			return i + 1
		}
	}
	return -1
}