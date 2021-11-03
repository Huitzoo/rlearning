package v1

import (
	"reinforcement/rfmodels/qlearning/components"
	"reinforcement/stages"
)

type QTable struct {
	states  []*components.State
	rewards []float64
	qTable  []float64
	maxX    int
	maxY    int
}

func (q *QTable) buildBadStatesAndActions(states []*components.State, badRewardValue float64) {
	actions := components.ValidateAction(states)
	for i, state := range states {
		q.states[state.ID] = state
		q.rewards[(components.TotalBasicActions*state.ID)+int(actions[i])] = badRewardValue
	}
}

func (q *QTable) completeRewardsValues(idx int, rewardValue float64) {
	for i := 0; i < components.TotalBasicActions; i++ {
		id := (components.TotalBasicActions * q.states[idx].ID) + i
		if q.rewards[id] == 0 {
			q.rewards[id] = rewardValue
		}
	}
}

func NewQBasicTable(stage stages.StageInterface) *QTable {
	sizeGrid := stage.GetSizeState()
	size := sizeGrid[0] * sizeGrid[1]

	qtable := &QTable{
		states:  make([]*components.State, size),
		rewards: make([]float64, size*components.TotalBasicActions),
		qTable:  make([]float64, size*components.TotalBasicActions),
		maxX:    sizeGrid[0],
		maxY:    sizeGrid[1],
	}

	badRewardValue := stage.GetBadRewardValue()
	rewardValue := stage.GetRewardValue()
	badActions := stage.GetBadAction()

	for {
		badAction := badActions()
		if badAction == nil {
			break
		}
		states := components.NewStates(badAction, sizeGrid[0])
		qtable.buildBadStatesAndActions(
			states,
			badRewardValue,
		)
	}
	row := 0
	column := 0

	for i := 0; i < size; i++ {
		if i%sizeGrid[0] == 0 && i != 0 {
			row++
			column = 0
		}
		if qtable.states[i] == nil {
			qtable.states[i] = components.NewState(row, column, sizeGrid[0])
		}
		column++
		qtable.completeRewardsValues(i, rewardValue)
	}
	return qtable
}

func (q *QTable) GetActionAndMaxScore(state int) (components.Action, float64) {
	idx := components.TotalBasicActions * state
	actions := q.qTable[idx : idx+components.TotalBasicActions]

	do, doAction := CalculateMaxOfArray(actions)

	return components.Action(doAction), do
}

func (q *QTable) getState(stateID int) *components.State {
	for i, state := range q.states {
		if state.ID == stateID {
			return q.states[i]
		}
	}
	return nil
}

func (q *QTable) Step(
	action components.Action,
	stateID int,
) (int, float64) {
	state := q.getState(stateID)
	if state == nil {
		return 0, 0
	}

	newCoords := action.OpareteStateWithAction(state.Coords)
	newStateID := q.maxX*newCoords.Y + newCoords.X

	if !newCoords.ValidateAroundCoords(q.maxX, q.maxY) {
		return state.ID, -100
	}

	idx := components.TotalBasicActions * stateID
	actions := q.rewards[idx : idx+components.TotalBasicActions]
	reward := actions[int(action)]

	return newStateID, reward
}

func (q *QTable) GetCurrentScoreAndReward(
	action components.Action,
	stateID int,
) (float64, float64) {
	idx := components.TotalBasicActions * stateID
	actions := q.qTable[idx : idx+components.TotalBasicActions]
	rewards := q.rewards[idx : idx+components.TotalBasicActions]
	return actions[int(action)], rewards[int(action)]
}

func (q *QTable) SetNewScore(stateID int, action components.Action, score float64) {
	idx := components.TotalBasicActions * stateID
	q.qTable[idx+int(action)] = score
}

func CalculateMaxOfArray(actions []float64) (float64, int) {

	var actionValue float64 = -9999999999999

	var doAction int

	for i, value := range actions {
		if value > actionValue {
			actionValue = value
			doAction = i
		}
	}

	return actionValue, doAction
}

func (q *QTable) GetTable() []float64 {
	return q.qTable
}

func (q *QTable) GetStates() []*components.State {
	return q.states
}
