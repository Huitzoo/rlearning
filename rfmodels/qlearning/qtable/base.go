package qtable

import (
	"reinforcement/stages"

	"reinforcement/rfmodels/qlearning/components"
)

type QTableInterface interface {
	Step(components.Action, []int) ([]int, float64)
	GetActionWithMaxScore([]int) (components.Action, float64)
	GetCurrentScore([]int, components.Action) float64
	SetNewScore([]int, components.Action, float64)
	PrintTable()
}

func NewQTable(env stages.StageInterface) QTableInterface {
	challenge := env.GetChallenge()
	switch challenge {
	case 1:
		return NewQBasicTable(env)
	case 2:
		return NewQTableWithObstables(env)
	default:
		return nil
	}
}
