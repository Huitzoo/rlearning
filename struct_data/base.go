package structdata

import "reinforcement/tools"

type TensorTag interface {
	SlideMatrix(ranges tools.RangesMatrix)
}

type Matrix struct {
	data []float64
	rTag int
}
