package main

func part2(lines []string) int {
	operators := []string{ADD_OPERATOR, MULTIPLY_OPERATOR, CONCAT_OPERATOR}

	calibrationSum := determineCalibrationSum(lines, operators)

	return calibrationSum
}
