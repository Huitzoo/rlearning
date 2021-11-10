package main

import (
	"fmt"
	"reinforcement/rfmodels/qlearning"
	"reinforcement/stages"
	"reinforcement/tools/views"
)

func main() {
	stage := stages.NewBasicStages("./stage2.yml")
	model := qlearning.NewQLearning(stage)

	if !model.LoadStage() {
		fmt.Println("Invalid challenge")
		return
	}

	model.Run()
	away, err := model.ValidateTable()
	if err != nil {
		fmt.Println("Agent cant learn")
		return
	}
	views.PrintQTable(away, stage)
}
