package main

import (
	"playbook-templates/structs"
	"playbook-templates/utils"
)

var args map[string]string

//params : source=pagerduty communication=slack data-collector=kubernetes
func main() {
	args := make(map[string]string)

	args["source"] = "datadog"
	args["communication"] = "slack"
	args["remediation"] = "kubernetes"
	template := LoadTemplate("", "master-template.yaml")

	GeneratedPlaybook(template, args, "playbook.yaml")
}

func GeneratedPlaybook(template structs.MasterTemplate, args map[string]string, playbookPath string) string {
	steps, inputs := GetSteps(template, args)

	playbook := structs.Playbook{
		Name:     template.Name,
		Steps:    steps,
		Inputs:   inputs,
		Triggers: GetTriggers(template, args),
	}

	utils.WriteYamlFile(&playbook, playbookPath)

	return playbookPath

}

func GetTriggers(template structs.MasterTemplate, a map[string]string) structs.Triggers {
	var webhooks []structs.Webhooks

	for _, trigger := range template.Triggers {
		impl := LoadImplementationFromYaml(trigger.Concrete, template.TemplateDirectory, args)
		webhooks = append(webhooks, impl.Triggers.Webhooks...)
	}

	return structs.Triggers{
		Webhooks: webhooks,
	}
}

func GetSteps(template structs.MasterTemplate, args map[string]string) ([]structs.Step, map[string]interface{}) {
	var (
		steps []structs.Step
	)
	inputs := make(map[string]interface{})

	for _, step := range template.Blocks {
		impl := LoadImplementationFromYaml(step.Concrete, template.TemplateDirectory, args)
		for _, block := range impl.Blocks {
			if block.Name == step.Name {
				steps = append(steps, structs.Step{Text: "# " + block.Name})
				steps = append(steps, block.Steps...)
				if len(block.Inputs) > 0 {
					for input := range block.Inputs {
						if _, ok := inputs[input]; !ok {
							inputs[input] = block.Inputs[input]
						}
					}
				}
			}
		}

	}
	return steps, inputs
}

func LoadTemplate(templateDirectoryPath string, templateFile string) structs.MasterTemplate {
	template := structs.MasterTemplate{}
	utils.LoadYamlFile(templateDirectoryPath+templateFile, &template)

	template.TemplateDirectory = templateDirectoryPath
	return template
}

func LoadImplementationFromYaml(concrete string, directoryPath string, args map[string]string) structs.Implementation {

	implem := structs.Implementation{}
	utils.LoadYamlFile(directoryPath+"implementation/"+concrete+"/"+args[concrete]+".yaml", &implem)

	return implem
}
