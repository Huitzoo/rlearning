package main

import (
	"fmt"
	"reinforcement/rfmodels/qlearning"
	"reinforcement/stages"
	"reinforcement/tools/views"
)

func startLearning() error {
	stage := stages.NewBasicStages("./stage3.yml")
	model := qlearning.NewQLearning(stage)

	if !model.LoadStage() {
		fmt.Println("Invalid challenge")
		return nil
	}

	model.Run()

	away, err := model.ValidateTable()
	if err != nil {
		fmt.Println("Agent cant learn")
		return err
	}
	views.PrintQTable(away, stage)
	return nil
}

func main() {
	for {
		if err := startLearning(); err == nil {
			break
		}
		//var second string
		//fmt.Scanln(&second)
	}

}
