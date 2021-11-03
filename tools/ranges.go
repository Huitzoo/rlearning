package tools

type RangesMatrix struct {
	ColumnRange []int
	RowRange    []int
}

type Coords struct {
	X int
	Y int
}

func (coords Coords) ValidateAroundCoords(maxX, maxY int) bool {
	if coords.X == -1 || coords.Y == -1 || coords.X == maxX || coords.Y == maxX {
		return false
	}
	return true
}

func ArrayIntsToCoords(data []int) Coords {
	return Coords{X: data[0], Y: data[1]}
}
