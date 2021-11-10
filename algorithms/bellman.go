package algorithms

func BellManEquation(
	currentScore,
	alpha,
	reward,
	discountFactor,
	nextMaxScore float64,
) float64 {
	return currentScore + alpha*(reward+(discountFactor*nextMaxScore)-currentScore)
}
