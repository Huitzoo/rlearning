package components

import (
	"math"
	"math/rand"
	"reinforcement/tools"
	"time"
)

type Action int

const TotalBasicActions int = 4

const (
	Up Action = iota
	Down
	Left
	Right
)

func (a Action) OpareteStateWithAction(coords tools.Coords) tools.Coords {
	newCoords := tools.Coords{}

	switch a {
	case Up:
		newCoords.X = coords.X
		newCoords.Y = coords.Y - 1
	case Down:
		newCoords.X = coords.X
		newCoords.Y = coords.Y + 1
	case Left:
		newCoords.X = coords.X - 1
		newCoords.Y = coords.Y
	default:
		newCoords.X = coords.X + 1
		newCoords.Y = coords.Y
	}
	return newCoords
}

func ValidateAction(states []*State) []Action {
	if math.Abs(float64(states[0].ID-states[1].ID)) == 1 {
		if states[0].ID > states[1].ID {
			return []Action{Left, Right}
		} else {
			return []Action{Right, Left}
		}
	} else {
		if states[0].ID > states[1].ID {
			return []Action{Up, Down}
		} else {
			return []Action{Down, Up}
		}
	}
}

func GetRandomAction() Action {
	rand.Seed(time.Now().UTC().UnixNano())
	do := 1 + rand.Int63n(3)
	return Action(do)
}
