package structs

type Step struct {
	Action string            `yaml:"action" json:"action"`
	Id     string            `yaml:"id" json:"id"`
	Name   string            `yaml:"name" json:"name"`
	Inputs map[string]string `yaml:"inputs,omitempty" json:"inputs"`
}

type Block struct {
	Name  string `yaml:"name" json:"name"`
	Steps []Step `yaml:"steps" json:"steps"`
}

type Implementation struct {
	Triggers Triggers `yaml:"triggers,omitempty"`
	Blocks   []Block  `yaml:"blocks" json:"blocks"`
}

type Triggers struct {
	Webhooks []Webhooks `yaml:"webhooks,omitempty" json:"webhooks"`
}

type Webhooks struct {
	Id     string            `yaml:"id" json:"id"`
	Name   string            `yaml:"name" json:"name"`
	Inputs map[string]string `yaml:"inputs,omitempty" json:"inputs"`
}
