package qlearning

import (
	"fmt"
	"math/rand"
	"reinforcement/rfmodels"
	"reinforcement/rfmodels/qlearning/components"
	"reinforcement/stages"
	structdata "reinforcement/struct_data"
	"reinforcement/tools"
	"time"
)

type QLearning struct {
	Stage       stages.StageInterface
	StageMatrix structdata.TensorTag
	QTable      *QTable
}

func NewQLearning(stage stages.StageInterface) rfmodels.RFModelInterface {
	model := &QLearning{}
	model.Stage = stage
	return model
}

func (q *QLearning) LoadStage() {
	q.QTable = NewQBasicTable(q.Stage)
}

func (q *QLearning) Run() {
	epochs := q.Stage.GetEpochs()
	initialState := q.Stage.GetInitialState()
	exploration := q.Stage.GetExploration()
	sizeGrid := q.Stage.GetSizeState()
	discountFactor := q.Stage.GetDiscountFactor()
	alpha := q.Stage.GetAlpha()
	goalStateID := q.Stage.GetGoalState()

	for i := 0; i < epochs; i++ {
		currentState := initialState

		fmt.Println("EPOCH: ", i)
		for {
			rand.Seed(time.Now().UTC().UnixNano())
			randomExploration := rand.Float64()
			var action components.Action

			if randomExploration > exploration {
				action = components.GetRandomAction()
			} else {
				action, _ = q.QTable.GetActionWithMaxScore(currentState)
			}

			nextState, reward, _ := q.QTable.Step(action, currentState)
			currentScore := q.QTable.GetCurrentScore(currentState, action)
			_, nextMaxScore := q.QTable.GetActionWithMaxScore(nextState)

			newScore := currentScore + alpha*(reward+(discountFactor*nextMaxScore)-currentScore)
			//fmt.Println(nextMaxScore, alpha, reward, nextMaxScore, discountFactor, currentScore, nextMaxScore)
			//fmt.Println(newScore)

			nextStateID := components.CalculateIDStateByCoords(
				tools.ArrayIntsToCoords(nextState), sizeGrid[0],
			)

			if nextStateID == goalStateID {
				q.QTable.SetNewScore(currentState, action, 10)
				break
			}
			q.QTable.SetNewScore(currentState, action, newScore)

			currentState = nextState
		}
	}
}

func (q *QLearning) GetResults() rfmodels.TablesInterface {
	return nil
}

func (q *QLearning) PrintTable() {
	state := q.Stage.GetInitialState()
	goalStateID := q.Stage.GetGoalState()
	coords := tools.ArrayIntsToCoords(state)
	sizeGrid := q.Stage.GetSizeState()

	currentStateID := components.CalculateIDStateByCoords(coords, sizeGrid[0])
	maxPrints := 10000
	counts := 0
	for currentStateID != goalStateID {
		if counts == maxPrints {
			fmt.Println("Error agent cant find way")
			break
		}
		coords := tools.ArrayIntsToCoords(state)
		stateID := components.CalculateIDStateByCoords(coords, sizeGrid[0])
		fmt.Println(state, q.QTable.table[stateID])
		action, _ := q.QTable.GetSecondBiggerAction(state)
		//next state calculation
		nextState, _, _ := q.QTable.Step(action, state)
		nextCoords := tools.ArrayIntsToCoords(nextState)
		//set state
		state = nextState
		currentStateID = components.CalculateIDStateByCoords(nextCoords, sizeGrid[0])
		counts++
	}

	fmt.Println("--------------------------------------------")
	rows := 0
	columns := 0
	for i := 0; i < len(q.QTable.table); i++ {

		if i%components.TotalBasicActions == 0 && i != 0 {
			rows++
			columns = 0
			fmt.Println()
		}
		fmt.Print("(", columns, rows, ")", q.QTable.table[i])
		columns++
	}
}
