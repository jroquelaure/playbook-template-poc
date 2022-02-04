package utils

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

func LoadYamlFile(file string, output interface{}) {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("yamlFile.Get err %v ", err)
	}

	err = yaml.Unmarshal(yamlFile, output)
	if err != nil {
		log.Printf("Unmarshal: %v", err)
	}
}

func WriteYamlFile(input interface{}, fileName string) {

	data, err := yaml.Marshal(input)
	if err != nil {
		log.Fatal(err)
	}

	err2 := ioutil.WriteFile(fileName, []byte(data), 0644)

	if err2 != nil {

		log.Fatal(err2)
	}

	fmt.Println("data written")
}
