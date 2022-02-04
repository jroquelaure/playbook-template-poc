package main

import (
	"os"
	"playbook-templates/utils"
	"testing"
)

var templateDirectory string
var templateFile string
var output string

func TestMain(m *testing.M) {

	args = make(map[string]string)

	args["source"] = "datadog"
	args["communication"] = "slack"
	args["remediation"] = "kubernetes"

	templateDirectory = "/Users/jon/workspace/lab/playbook-templates/templates/IncidentResponse/"
	templateFile = "master.yaml"
	output = "datadog-kube-slack.yaml"
	exitVal := m.Run()

	args["source"] = "pagerduty"
	output = "pagerduty-kube-slack.yaml"
	exitVal = m.Run()

	//main()
	os.Exit(exitVal)
}

func TestLoadTemplate(t *testing.T) {
	template := LoadTemplate(templateDirectory, templateFile)

	t.Logf(template.Name)

	t.Logf("success")
}

func TestLoadImplementationFromYaml(t *testing.T) {

	imp := LoadImplementationFromYaml("source", templateDirectory, args)

	for a, _ := range imp.Blocks {
		t.Logf("d%dfzf ", a)
	}

	t.Logf("success")

}

func TestGetSteps(t *testing.T) {
	template := LoadTemplate(templateDirectory, templateFile)

	steps := GetSteps(template, args)

	for a, _ := range steps {
		t.Logf("d%dfzf ", a)
	}

	t.Logf("success")

}

func TestGeneratedPlaybook(t *testing.T) {

	template := LoadTemplate(templateDirectory, templateFile)

	file := GeneratedPlaybook(template, args, output)

	t.Logf("d%sfzf ", file)

	result := make(map[string]interface{})

	utils.LoadYamlFile(file, &result)
	if result["name"] != template.Name {
		t.Errorf("Expected name %s but was %s", template.Name, result["name"])
	}
	if len(result["steps"].([]interface{})) == 0 {
		t.Errorf("No steps")
	}
	triggers := result["triggers"].(map[string]interface{})
	if len(triggers["webhooks"].([]interface{})) == 0 {
		t.Errorf("No triggers for %s ", file)
	}
	t.Logf("success")

}
