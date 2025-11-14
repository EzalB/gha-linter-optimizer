package parser

import (
	"os"
	"gopkg.in/yaml.v3"
)

type Workflow struct {
	Jobs map[string]Job `yaml:"jobs"`
	Path string          `yaml:"-"`
}

type Job struct {
	Name     string            `yaml:"name"`
	RunsOn   string            `yaml:"runs-on"`
	Steps    []Step            `yaml:"steps"`
	Line     int               `yaml:"-"` // line number in YAML
	Strategy map[string]any    `yaml:"strategy,omitempty"`
	Env      map[string]string `yaml:"env,omitempty"`
}

type Step struct {
	Name string            `yaml:"name"`
	Uses string            `yaml:"uses"`
	Run  string            `yaml:"run"`
	If   string            `yaml:"if,omitempty"`
	Env  map[string]string `yaml:"env,omitempty"`
	Line int               `yaml:"-"` // line number in YAML
}

func ParseWorkflow(path string) (*Workflow, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var root yaml.Node
	if err := yaml.Unmarshal(data, &root); err != nil {
		return nil, err
	}

	wf := &Workflow{Path: path, Jobs: make(map[string]Job)}

	// find the 'jobs:' node
	for i := 0; i < len(root.Content[0].Content); i += 2 {
		keyNode := root.Content[0].Content[i]
		valNode := root.Content[0].Content[i+1]

		if keyNode.Value == "jobs" && valNode.Kind == yaml.MappingNode {
			for j := 0; j < len(valNode.Content); j += 2 {
				jobKey := valNode.Content[j]
				jobVal := valNode.Content[j+1]

				job := Job{
					Name:   extractString(jobVal, "name"),
					RunsOn: extractString(jobVal, "runs-on"),
					Line:   jobKey.Line,
				}

				// extract steps
				for k := 0; k < len(jobVal.Content); k += 2 {
					key := jobVal.Content[k]
					val := jobVal.Content[k+1]
					if key.Value == "steps" && val.Kind == yaml.SequenceNode {
						for _, stepNode := range val.Content {
							step := Step{
								Line: stepNode.Line,
							}
							for x := 0; x < len(stepNode.Content); x += 2 {
								stepKey := stepNode.Content[x]
								stepVal := stepNode.Content[x+1]
								switch stepKey.Value {
								case "name":
									step.Name = stepVal.Value
								case "uses":
									step.Uses = stepVal.Value
								case "run":
									step.Run = stepVal.Value
								case "if":
									step.If = stepVal.Value
								case "env":
									env := make(map[string]string)
									for e := 0; e < len(stepVal.Content); e += 2 {
										envKey := stepVal.Content[e]
										envVal := stepVal.Content[e+1]
										env[envKey.Value] = envVal.Value
									}
									step.Env = env
								}
							}
							job.Steps = append(job.Steps, step)
						}
					}
				}
				wf.Jobs[jobKey.Value] = job
			}
		}
	}
	return wf, nil
}

func extractString(node *yaml.Node, key string) string {
	for i := 0; i < len(node.Content); i += 2 {
		k := node.Content[i]
		v := node.Content[i+1]
		if k.Value == key {
			return v.Value
		}
	}
	return ""
}