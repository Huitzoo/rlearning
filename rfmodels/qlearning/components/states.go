package components

import (
	"reinforcement/stages"
	"reinforcement/tools"
)

type State struct {
	Coords tools.Coords
	ID     int
}

func (s *State) CalculatePosition(columns int) {
	s.ID = columns*s.Coords.Y + s.Coords.X
}

func NewStates(states [][]int, columns int) []*State {
	coords := getStatesPosition(states)
	realStates := make([]*State, 2)
	stateOne := State{Coords: coords[0]}
	stateTwo := State{Coords: coords[1]}
	stateOne.CalculatePosition(columns)
	stateTwo.CalculatePosition(columns)
	realStates[0] = &stateOne
	realStates[1] = &stateTwo
	return realStates
}

func NewState(row, column, numberOfColumns int) *State {
	state := &State{Coords: tools.Coords{
		X: column, Y: row,
	}}
	state.CalculatePosition(numberOfColumns)
	return state
}

func getStatesPosition(states stages.StatesPosition) []tools.Coords {
	return []tools.Coords{
		{
			X: states[0][0],
			Y: states[0][1],
		}, {
			X: states[1][0],
			Y: states[1][1],
		},
	}
}

func CalculateIDStateByCoords(coords tools.Coords, columns int) int {
	return columns*coords.Y + coords.X
}
