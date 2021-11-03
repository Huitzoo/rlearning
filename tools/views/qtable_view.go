package views

import (
	"fmt"
	"reinforcement/rfmodels"
	"reinforcement/rfmodels/qlearning/components"
	"reinforcement/stages"
)

func PrintQTable(q rfmodels.TablesInterface, stage stages.StageInterface) {

	initStateCoords := stage.GetInitialState()
	size := stage.GetSizeState()
	nCols := size[0]
	stateID := nCols*initStateCoords[1] + initStateCoords[0]
	table := q.GetTable()
	states := q.GetStates()

	//goalState := stage.GetGoalState()

	for i := 0; i < len(table); i++ {
		newStateID := nCols * i

		fmt.Println(
			table[idx:idx+components.TotalBasicActions], action,
		)
		fmt.Println(
			states[stateID].Coords, "->", states[newStateID].Coords,
		)

		stateID = newStateID
	}

	//gocv.WaitKey(0)
}

func getMax(table []float64) components.Action {
	action := 0
	valueMin := -99999999999.0

	for doAction, value := range table {
		if value > valueMin {
			valueMin = value
			action = doAction
		}
	}
	return components.Action(action)
}
