package main

func part1(lines []string) int {
	operators := []string{ADD_OPERATOR, MULTIPLY_OPERATOR}

	calibrationSum := determineCalibrationSum(lines, operators)

	return calibrationSum
}
