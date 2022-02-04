package structs

type MasterTemplate struct {
	Name              string         `yaml:"name" json:"name"`
	Type              string         `yaml:"type" json:"type"`
	Desc              string         `yaml:"desc" json:"desc"`
	Triggers          []BaseTriggers `yaml:"triggers" json:"triggers"`
	Blocks            []BaseBlock    `yaml:"blocks" json:"blocks"`
	TemplateDirectory string         `yaml:"templateDirectory" json:"templateDirectory"`
}

type BaseBlock struct {
	Concrete string `yaml:"concrete" json:"concrete"`
	Name     string `yaml:"name" json:"name"`
}

type BaseTriggers struct {
	Concrete string `yaml:"concrete" json:"concrete"`
}
