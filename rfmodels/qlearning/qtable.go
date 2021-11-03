package qlearning

import (
	"reinforcement/rfmodels/qlearning/components"
	"reinforcement/stages"
	"reinforcement/tools"
)

type QTable struct {
	columns, rows int
	reward        float64
	badReward     float64
	badRewards    map[int][]bool
	table         [][]float64
}

func NewQBasicTable(stage stages.StageInterface) *QTable {
	sizeGrid := stage.GetSizeState()
	columns := sizeGrid[0]
	rows := sizeGrid[1]

	qtable := &QTable{
		badRewards: make(map[int][]bool),
		table:      make([][]float64, rows*columns),
	}

	badActions := stage.GetBadAction()

	for {
		badAction := badActions()
		if badAction == nil {
			break
		}
		states := components.NewStates(badAction, columns)
		qtable.setStatesWithBadRewards(
			states,
		)
	}
	qtable.columns = columns
	qtable.rows = rows
	qtable.reward = stage.GetRewardValue()
	qtable.badReward = stage.GetBadRewardValue()

	qtable.SetUpTable()

	return qtable
}

func (q *QTable) setStatesWithBadRewards(states []*components.State) {
	actions := components.ValidateAction(states)

	for i, state := range states {
		if _, exits := q.badRewards[state.ID]; !exits {
			q.badRewards[state.ID] = make([]bool, components.TotalBasicActions)
		}
		q.badRewards[state.ID][actions[i]] = true
	}
}

func (q *QTable) SetUpTable() {
	for i := range q.table {
		q.table[i] = make([]float64, components.TotalBasicActions)
	}
}

func (q *QTable) Step(
	action components.Action,
	state []int,
) ([]int, float64, bool) {
	coords := tools.ArrayIntsToCoords(state)
	newCoords := action.OpareteStateWithAction(coords)
	badState := false
	if !newCoords.ValidateAroundCoords(q.columns, q.rows) {
		badState = true
		return state, q.badReward, badState
	}

	reward := q.reward

	newState := []int{newCoords.X, newCoords.Y}
	stateID := components.CalculateIDStateByCoords(coords, q.columns)

	if actions, exits := q.badRewards[stateID]; exits {
		if actions[action] {
			reward = q.badReward
		}
	}

	return newState, reward, badState
}
func (q *QTable) GetActionWithMaxScore(state []int) (components.Action, float64) {
	coords := tools.ArrayIntsToCoords(state)
	stateID := components.CalculateIDStateByCoords(coords, q.columns)
	stateActions := q.table[stateID]
	score := stateActions[0]
	action := 0
	for i := 1; i < components.TotalBasicActions; i++ {
		if score <= stateActions[i] {
			score = stateActions[i]
			action = i
		}
	}

	if score == 0 {
		score = 0.1
	}

	return components.Action(action), score
}

func (q *QTable) GetCurrentScore(state []int, action components.Action) float64 {
	coords := tools.ArrayIntsToCoords(state)
	stateID := components.CalculateIDStateByCoords(coords, q.columns)

	return q.table[stateID][action]
}

func (q *QTable) SetNewScore(state []int, action components.Action, score float64) {
	coords := tools.ArrayIntsToCoords(state)
	stateID := components.CalculateIDStateByCoords(coords, q.columns)

	q.table[stateID][action] += score / 100
}

func (q *QTable) GetSecondBiggerAction(state []int) (components.Action, float64) {
	coords := tools.ArrayIntsToCoords(state)
	stateID := components.CalculateIDStateByCoords(coords, q.columns)
	stateActions := q.table[stateID]
	score := stateActions[0]
	action := 0

	if score == q.reward {
		score = -1
	}

	for i := 1; i < components.TotalBasicActions; i++ {
		if score <= stateActions[i] {
			score = stateActions[i]
			action = i
		}
	}
	return components.Action(action), score
}
