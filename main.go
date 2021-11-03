package main

import (
	"reinforcement/rfmodels/qlearning"
	"reinforcement/stages"
)

func main() {
	stage := stages.NewBasicStages("./stage1.yml")
	model := qlearning.NewQLearning(stage)

	model.LoadStage()
	model.Run()
	model.PrintTable()
	//views.PrintQTable(model.GetResults(), stage)
}
