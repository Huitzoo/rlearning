package main

import (
	"fmt"
	"os"
	"reinforcement/rfmodels/qlearning"
	"reinforcement/stages"
	"reinforcement/tools/views"
)

func startLearning() error {
	if len(os.Args) < 2 {
		fmt.Println("You have to put type of challenge and yml stage")
		return nil
	}

	challenge := os.Args[1]
	path := os.Args[2]

	stage := stages.NewBasicStages(path, challenge)
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
	}

}
