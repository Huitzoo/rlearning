package rfmodels

import (
	"reinforcement/rfmodels/qlearning/components"
)

type RFModelInterface interface {
	LoadStage()
	Run()
	GetResults() TablesInterface
	PrintTable()
}

type TablesInterface interface {
	GetTable() []float64
	GetStates() []*components.State
	Step(components.Action, int) (int, float64)
}
