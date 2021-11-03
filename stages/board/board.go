package board

import (
	"image"
	"image/color"

	"gocv.io/x/gocv"
)

var (
	blue = color.RGBA{0, 0, 255, 0}
)

type BoardInterface interface {
	CreateBoard(int, int, string)
	PaintPoint(int, int)
	ClearBoard()
}

type Board struct {
	board gocv.Mat
	x, y  int
	name  string
}

func (b *Board) CreateBoard(x, y int, name string) {
	b.board = gocv.NewMatWithSize(x, y, 1)
	b.name = name
	b.x = x
	b.y = y
}

func (b *Board) PaintPoint(x, y int) {
	pt := image.Point{
		X: x,
		Y: y,
	}
	gocv.Line(
		&b.board,
		pt, pt,
		blue,
		1,
	)
}

func (b *Board) ClearBoard() {
	b.board = gocv.NewMatWithSize(b.x, b.y, 1)
}

func NewBoard(x, y int) BoardInterface {
	board := &Board{}
	board.CreateBoard(x, y, "Board ")
	return board
}
