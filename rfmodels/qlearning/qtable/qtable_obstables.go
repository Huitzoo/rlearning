package qtable

import (
	"fmt"
	"math"
	"reinforcement/rfmodels/qlearning/components"
	"reinforcement/stages"
	"reinforcement/tools"
)

type QTableWithObstables struct {
	columns, rows int
	reward        float64
	badReward     float64
	badRewards    map[int]struct{}
	table         [][]float64
}

func NewQTableWithObstables(stage stages.StageInterface) QTableInterface {
	sizeGrid := stage.GetSizeState()
	columns := sizeGrid[0]
	rows := sizeGrid[1]

	QTableWithObstables := &QTableWithObstables{
		badRewards: make(map[int]struct{}),
		table:      make([][]float64, rows*columns),
	}

	obstacles := stage.GetBadAction()

	for {
		obstacle := obstacles()
		state, ok := obstacle.([]int)
		if !ok {
			break
		}
		QTableWithObstables.setStateWithBadRewards(
			components.NewState(state[0], state[1], columns),
		)
	}

	QTableWithObstables.columns = columns
	QTableWithObstables.rows = rows
	QTableWithObstables.reward = stage.GetRewardValue()
	QTableWithObstables.badReward = stage.GetBadRewardValue()

	QTableWithObstables.SetUpTable()

	return QTableWithObstables
}

func (q *QTableWithObstables) setStateWithBadRewards(state *components.State) {
	q.badRewards[state.ID] = struct{}{}
}

func (q *QTableWithObstables) SetUpTable() {
	for i := range q.table {
		q.table[i] = make([]float64, components.TotalBasicActions)
	}
}

func (q *QTableWithObstables) Step(
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

	if _, exits := q.badRewards[stateID]; exits {
		reward = q.badReward
	}

	return newState, reward, badState
}

func (q *QTableWithObstables) GetActionWithMaxScore(state []int) (components.Action, float64) {
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

func (q *QTableWithObstables) GetCurrentScore(state []int, action components.Action) float64 {
	coords := tools.ArrayIntsToCoords(state)
	stateID := components.CalculateIDStateByCoords(coords, q.columns)

	return q.table[stateID][action]
}

func (q *QTableWithObstables) SetNewScore(state []int, action components.Action, score float64) {
	coords := tools.ArrayIntsToCoords(state)
	stateID := components.CalculateIDStateByCoords(coords, q.columns)

	q.table[stateID][action] += score / 1000
}

func (q *QTableWithObstables) GetSecondBiggerAction(state []int) (components.Action, float64) {
	coords := tools.ArrayIntsToCoords(state)
	stateID := components.CalculateIDStateByCoords(coords, q.columns)
	stateActions := q.table[stateID]
	score := stateActions[0]
	action := 0

	if score == q.reward {
		score = -1
	}

	for i := 1; i < components.TotalBasicActions; i++ {
		if score <= stateActions[i] && !math.IsNaN(stateActions[i]) {
			score = stateActions[i]
			action = i
		}
	}
	return components.Action(action), score
}

func (q *QTableWithObstables) PrintTable() {
	rows := 0
	columns := 0
	for i := 0; i < len(q.table); i++ {

		if i%components.TotalBasicActions == 0 && i != 0 {
			rows++
			columns = 0
			fmt.Println()
		}
		fmt.Print("(", columns, rows, ")", q.table[i])
		columns++
	}
}
