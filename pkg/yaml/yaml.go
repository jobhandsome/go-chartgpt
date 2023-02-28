package yaml

import (
	"go-chatgpt/app/frontend/v1/config"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func Init() {
	getwd, _ := os.Getwd()
	fileName := getwd + "/etc/frontend.yaml"

	config.Config = new(config.FrontendYaml)

	yamlFile, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("read file err %v\\n", err)
		return
	}

	err = yaml.Unmarshal(yamlFile, config.Config)
	if err != nil {
		log.Fatalf("yaml unmarshal: %v\n", err)
		return
	}
}
