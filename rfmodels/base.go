package rfmodels

import (
	"reinforcement/tools"
)

type RFModelInterface interface {
	Run()
	LoadStage() bool
	ValidateTable() ([]tools.Coords, error)
}
