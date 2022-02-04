package structs

type Playbook struct {
	Name     string   `yaml:"name" json:"name"`
	Triggers Triggers `yaml:"triggers" json:"triggers"`
	Steps    []Step   `yaml:"steps" json:"steps"`
}
