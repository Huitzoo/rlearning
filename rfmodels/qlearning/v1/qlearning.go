package v1

import (
	"fmt"
	"math/rand"
	"reinforcement/rfmodels"
	"reinforcement/rfmodels/qlearning/components"
	"reinforcement/stages"
	structdata "reinforcement/struct_data"
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

func (q *QLearning) GetResults() rfmodels.TablesInterface {
	return q.QTable
}

func (q *QLearning) Run() {
	epochs := q.Stage.GetEpochs()
	exploration := q.Stage.GetExploration()
	sizeGrid := q.Stage.GetSizeState()
	initialState := q.Stage.GetInitialState()
	InitialStateID := sizeGrid[0]*initialState[0] + initialState[1]
	discountFactor := q.Stage.GetDiscountFactor()
	alpha := q.Stage.GetAlpha()
	goalState := q.Stage.GetGoalState()
	//punishment := q.Stage.GetBadRewardValue()

	for i := 0; i < epochs; i++ {
		currentState := InitialStateID
		steps := q.Stage.GetSteps()
		fmt.Println("EPOCH: ", i)
		for {
			rand.Seed(time.Now().UTC().UnixNano())

			randomExploration := rand.Float64()
			var action components.Action

			if randomExploration > exploration {
				action = components.GetRandomAction()
			} else {
				action, _ = q.QTable.GetActionAndMaxScore(InitialStateID)
			}

			nextState, reward := q.QTable.Step(action, currentState)

			currentScore, _ := q.QTable.GetCurrentScoreAndReward(action, currentState)
			_, nextMaxScore := q.QTable.GetActionAndMaxScore(nextState)

			newScore := currentScore + alpha*(reward+(discountFactor*nextMaxScore))

			q.QTable.SetNewScore(currentState, action, newScore)

			if nextState == goalState {
				break
			}

			currentState = nextState
			steps--
		}
	}
}

func (q *QLearning) PrintTable() {
	fmt.Println(q.QTable.qTable)
}
