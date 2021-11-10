package stages

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

func NewBasicStages(path, challenge string) StageInterface {
	stage := FactoryStages(challenge)
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

func FactoryStages(challenge string) StageInterface {
	switch challenge {
	case "obstacles":
		return &ObstacleStage{}
	case "movement":
		return &BasicStages{}
	default:
		return &ObstacleStage{}
	}
}
