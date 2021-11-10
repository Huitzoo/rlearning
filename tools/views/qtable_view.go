package views

import (
	"reinforcement/stages"
	"reinforcement/stages/board"
	"reinforcement/tools"
	"time"

	"gocv.io/x/gocv"
)

func PrintQTable(away []tools.Coords, stage stages.StageInterface) {
	window := gocv.NewWindow("Hello")
	for {
		for _, coords := range away {
			boardStage := stage.GetBoard()
			boardStage.PaintPoint(coords.X, coords.Y, board.BLUE)
			time.Sleep(400 * time.Millisecond)
			window.IMShow(boardStage.ReturnImageBoard())
			if window.WaitKey(1) >= 0 {
				break
			}
		}
	}
}
