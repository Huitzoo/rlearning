package structdata

import (
	"errors"
	"reinforcement/tools"
)

type Matrix2D struct {
	Matrices []Matrix
}

func NewMatrix2D(r, c int, v []float64) (*Matrix2D, error) {
	if len(v)%(r*c) != 0 {
		return nil, errors.New("cant convert into 2d matrix")
	}

	getSliceOfArray := func(start, end int, arr []float64) []float64 {
		return arr[start:end]
	}
	m := &Matrix2D{}
	rowsCount := 0
	for i := 0; i < len(v); i += c {
		smallMatrix := getSliceOfArray(i, c+i, v)
		m.Matrices = append(
			m.Matrices,
			Matrix{smallMatrix, rowsCount},
		)
		rowsCount++
	}
	return m, nil
}

func (m *Matrix2D) SlideMatrix(ranges tools.RangesMatrix) *Matrix2D {
	resultMatrix := &Matrix2D{}
	rowsCount := 0
	for _, matrix := range m.Matrices {
		if matrix.rTag >= ranges.RowRange[0] && matrix.rTag < ranges.RowRange[1] {
			resultMatrix.Matrices = append(
				resultMatrix.Matrices,
				Matrix{matrix.data[ranges.ColumnRange[0]:ranges.ColumnRange[1]], rowsCount},
			)
			rowsCount++
		}
	}
	return resultMatrix
}

/*
x,y
cRange(0,1),yRange(0,3)
:,:
*/
