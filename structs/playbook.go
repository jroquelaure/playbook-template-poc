package structs

type Playbook struct {
	Name     string                 `yaml:"name" json:"name"`
	Triggers Triggers               `yaml:"triggers" json:"triggers"`
	Steps    []Step                 `yaml:"steps" json:"steps"`
	Inputs   map[string]interface{} `yaml:"inputs,omitempty" json:"inputs"`
}
