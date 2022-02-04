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
	playbook := structs.Playbook{
		Name:     template.Name,
		Steps:    GetSteps(template, args),
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

func GetSteps(template structs.MasterTemplate, args map[string]string) []structs.Step {
	var (
		steps []structs.Step
	)
	for _, step := range template.Blocks {
		impl := LoadImplementationFromYaml(step.Concrete, template.TemplateDirectory, args)
		for _, block := range impl.Blocks {
			if block.Name == step.Name {
				steps = append(steps, block.Steps...)
			}
		}

	}
	return steps
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
