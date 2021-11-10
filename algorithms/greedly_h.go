package algorithms

import "math/rand"

const GreddlyZero = 0.4

var realEpochs = 0.0

func UpdateGreedlyByFactors(
	epochs int,
	currentEpoch int,
	greedlyExploration float64,
	factor float64,
) float64 {

	if float64(currentEpoch/epochs) > GreddlyZero {
		return 0.1
	}
	return greedlyExploration - factor
}

func UpdateGreedlyBySteps(
	stepFactor float64,
) float64 {
	var exploration = 0.1
	if stepFactor > 0.2 && 0.8 <= stepFactor {
		exploration = (rand.Float64() * (0.7 - 0.4)) + 0.4
	} else if stepFactor > 0.8 {
		exploration = (rand.Float64() * (0.9 - 0.7)) + 0.7
	}
	return exploration
}

func CalculateFactor(epoch int, greedlyExploration float64) float64 {
	realEpochs = (float64(epoch) * greedlyExploration)
	return greedlyExploration / realEpochs
}
