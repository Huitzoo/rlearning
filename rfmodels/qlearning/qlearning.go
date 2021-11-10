package qlearning

import (
	"errors"
	"math/rand"
	"reinforcement/algorithms"
	"reinforcement/rfmodels"

	"reinforcement/rfmodels/qlearning/components"
	"reinforcement/rfmodels/qlearning/qtable"
	"reinforcement/stages"
	structdata "reinforcement/struct_data"
	"reinforcement/tools"
)

type QLearning struct {
	Stage       stages.StageInterface
	StageMatrix structdata.TensorTag
	QTable      qtable.QTableInterface
}

func NewQLearning(stage stages.StageInterface) rfmodels.RFModelInterface {
	model := &QLearning{}
	model.Stage = stage
	return model
}

func (q *QLearning) LoadStage() bool {
	qtable := qtable.NewQTable(q.Stage)
	if qtable != nil {
		q.QTable = qtable
		return true
	}
	return false
}

func (q *QLearning) Run() {
	epochs := q.Stage.GetEpochs()
	initialState := q.Stage.GetInitialState()
	greedlyExploration := q.Stage.GetExploration()
	sizeGrid := q.Stage.GetSizeState()
	discountFactor := q.Stage.GetDiscountFactor()
	alpha := q.Stage.GetAlpha()
	goalStateID := q.Stage.GetGoalState()
	currentSteps := 0
	_ = currentSteps
	tools.SetSeed()

	for i := 0; i < epochs; i++ {
		currentState := initialState
		steps := q.Stage.GetSteps()
		//fmt.Println("EPOCH: ", i)
		if i != 0 {
			stepFactor := float64(currentSteps) / float64(steps)
			greedlyExploration = algorithms.UpdateGreedlyBySteps(stepFactor)
			//fmt.Println("EPOCH: ", i, "greedlyExploration: ", greedlyExploration, "currentSteps: ", currentSteps, "stepFactor: ", stepFactor)
		}

		currentSteps = 0
		for steps != 0 {

			maxExploration := rand.Float64()
			var action components.Action

			if greedlyExploration > maxExploration {
				action = components.GetRandomAction()
			} else {
				action, _ = q.QTable.GetActionWithMaxScore(currentState)
			}

			nextState, reward := q.QTable.Step(action, currentState)
			currentScore := q.QTable.GetCurrentScore(currentState, action)
			_, nextMaxScore := q.QTable.GetActionWithMaxScore(nextState)

			newScore := algorithms.BellManEquation(
				currentScore, alpha, reward, discountFactor, nextMaxScore,
			)

			nextStateID := components.CalculateIDStateByCoords(
				tools.ArrayIntsToCoords(nextState), sizeGrid[0],
			)
			//fmt.Println("currentState: ", currentState, "action: ", action, "newScore: ", newScore, "nextMaxScore: ", nextMaxScore, "nextState: ", nextState)

			if nextStateID == goalStateID {
				q.QTable.SetNewScore(currentState, action, 10)
				break
			}
			q.QTable.SetNewScore(currentState, action, newScore)

			currentState = nextState
			steps--
			currentSteps++
		}
	}
	//q.QTable.PrintTable()
}

func (q *QLearning) ValidateTable() ([]tools.Coords, error) {
	state := q.Stage.GetInitialState()
	goalStateID := q.Stage.GetGoalState()
	coords := tools.ArrayIntsToCoords(state)
	sizeGrid := q.Stage.GetSizeState()
	currentStateID := components.CalculateIDStateByCoords(coords, sizeGrid[0])
	maxPrints := 500
	currentWay := []tools.Coords{
		coords,
	}

	for currentStateID != goalStateID {
		if maxPrints == 0 {
			return nil, errors.New("agent didn't learn the correct away")

		}
		action, _ := q.QTable.GetActionWithMaxScore(state)

		//next state calculation
		nextState, _ := q.QTable.Step(action, state)
		nextCoords := tools.ArrayIntsToCoords(nextState)
		currentWay = append(currentWay, nextCoords)
		//set state
		state = nextState
		currentStateID = components.CalculateIDStateByCoords(nextCoords, sizeGrid[0])
		maxPrints--
	}

	q.QTable.PrintTable()
	return currentWay, nil
}
