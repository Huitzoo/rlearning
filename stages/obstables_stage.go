package stages

import (
	"fmt"
	"math/rand"
	"reinforcement/stages/board"
)

type ObstacleStage struct {
	Size           []int   `yaml:"size"`
	InitialState   []int   `yaml:"initial_state"`
	Goal           []int   `yaml:"goal"`
	Danger         [][]int `yaml:"danger"`
	Epochs         int     `yaml:"epochs"`
	Steps          int     `yaml:"steps"`
	Exploration    float64 `yaml:"exploration"`
	Reward         float64 `yaml:"reward"`
	Punishment     float64 `yaml:"punishment"`
	Alpha          float64 `yaml:"alpha"`
	DiscountFactor float64 `yaml:"discount_factor"`
	Challenge      int64   `yaml:"challenge"`
}

func (s *ObstacleStage) GetBadAction() DangerActionStage {
	var num = len(s.Danger)
	return func() interface{} {
		num = num - 1
		if num != -1 {
			return s.Danger[num]
		}
		return nil
	}
}

func (s *ObstacleStage) GetBadRewardValue() float64 {
	return s.Punishment
}

func (s *ObstacleStage) GetRewardValue() float64 {
	return s.Reward
}

func (s *ObstacleStage) GetSizeState() SizeStage {
	return s.Size
}
func (s *ObstacleStage) GetSteps() int {
	return s.Steps
}
func (s *ObstacleStage) GetEpochs() int {
	return s.Epochs
}
func (s *ObstacleStage) GetExploration() float64 {
	return s.Exploration
}

func (s *ObstacleStage) GetAlpha() float64 {
	return s.Alpha
}

func (s *ObstacleStage) GetDiscountFactor() float64 {
	return s.DiscountFactor
}

func (s *ObstacleStage) GetGoalState() int {
	return s.Size[0]*s.Goal[1] + s.Goal[0]
}

func (s *ObstacleStage) GetChallenge() int64 {
	return s.Challenge
}

func (s *ObstacleStage) GetInitialState() []int {
	if len(s.InitialState) == 0 {
		return []int{rand.Intn(s.Size[0]), rand.Intn(s.Size[0])}
	} else {
		return s.InitialState
	}
}

func (s *ObstacleStage) GetBoard() board.BoardInterface {
	boardStage := board.NewBoard(s.Size[0], s.Size[1])

	obstacles := s.GetBadAction()

	for {
		obstacle := obstacles()
		state, ok := obstacle.([]int)
		fmt.Println(obstacle)
		if !ok {
			break
		}
		boardStage.PaintPoint(
			state[0], state[1], board.ORANGE,
		)
	}
	boardStage.BackUpInitialBoard()
	return boardStage
}
