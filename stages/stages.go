package stages

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

func NewBasicStages(path string) StageInterface {
	stage := &BasicStages{}
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, stage)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return stage
}
