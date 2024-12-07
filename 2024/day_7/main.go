package main

import (
	"fmt"
)

var operatorCombinations []string

const (
	MULTIPLY_OPERATOR = "*"
	ADD_OPERATOR      = "+"
	CONCAT_OPERATOR   = "|"
)

func main() {
	file := "input.txt"

	lines, err := ParseInput(file)
	if err != nil {
		message := fmt.Sprintf("%v", err)
		panic(message)
	}

	partOneCalibrationSum := part1(lines)
	partTwoCalibrationSum := part2(lines)

	fmt.Printf("Calibration sum (part 1) = %d\n", partOneCalibrationSum)
	fmt.Printf("Calibration sum (part 2) = %d\n", partTwoCalibrationSum)
}
