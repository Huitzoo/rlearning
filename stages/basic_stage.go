package stages

import (
	"math/rand"
	"reinforcement/stages/board"
)

type BasicStages struct {
	Size           []int     `yaml:"size"`
	InitialState   []int     `yaml:"initial_state"`
	Goal           []int     `yaml:"goal"`
	Danger         [][][]int `yaml:"danger"`
	Epochs         int       `yaml:"epochs"`
	Steps          int       `yaml:"steps"`
	Exploration    float64   `yaml:"exploration"`
	Reward         float64   `yaml:"reward"`
	Punishment     float64   `yaml:"punishment"`
	Alpha          float64   `yaml:"alpha"`
	DiscountFactor float64   `yaml:"discount_factor"`
	Challenge      int64     `yaml:"challenge"`
}

func (s *BasicStages) GetBadAction() DangerActionStage {
	var num = len(s.Danger)
	return func() interface{} {
		num = num - 1
		if num != -1 {
			return s.Danger[num]
		}
		return nil
	}
}

func (s *BasicStages) GetBadRewardValue() float64 {
	return s.Punishment
}

func (s *BasicStages) GetRewardValue() float64 {
	return s.Reward
}

func (s *BasicStages) GetSizeState() SizeStage {
	return s.Size
}
func (s *BasicStages) GetSteps() int {
	return s.Steps
}
func (s *BasicStages) GetEpochs() int {
	return s.Epochs
}
func (s *BasicStages) GetExploration() float64 {
	return s.Exploration
}

func (s *BasicStages) GetAlpha() float64 {
	return s.Alpha
}

func (s *BasicStages) GetDiscountFactor() float64 {
	return s.DiscountFactor
}

func (s *BasicStages) GetGoalState() int {
	return s.Size[0]*s.Goal[1] + s.Goal[0]
}

func (s *BasicStages) GetChallenge() int64 {
	return s.Challenge
}

func (s *BasicStages) GetInitialState() []int {
	if len(s.InitialState) == 0 {
		return []int{rand.Intn(s.Size[0]), rand.Intn(s.Size[0])}
	} else {
		return s.InitialState
	}
}

func (s *BasicStages) GetBoard() board.BoardInterface {
	boardStage := board.NewBoard(s.Size[0], s.Size[1])
	boardStage.BackUpInitialBoard()

	return boardStage
}
