package stages

import (
	"reinforcement/stages/board"
)

type BadRewardStagePositions [][]int
type DangerActionStage func() BadRewardStagePositions
type SizeStage []int

type StageInterface interface {
	GetBadAction() DangerActionStage
	GetSizeState() SizeStage

	GetBadRewardValue() float64
	GetDiscountFactor() float64
	GetRewardValue() float64
	GetExploration() float64
	GetAlpha() float64

	GetSteps() int
	GetEpochs() int

	GetInitialState() []int
	GetGoalState() int
	GetBoard() board.BoardInterface
}

type StatesPosition BadRewardStagePositions
