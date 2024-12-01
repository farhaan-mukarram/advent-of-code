package main

import (
	"fmt"
	"math"
	"slices"
)

func part1(lines []string) int {
	var leftNums, rightNums []int
	var leftNum, rightNum int

	// Populate left and right arrays with numbers
	for _, line := range lines {
		fmt.Sscanf(line, "%d %d", &leftNum, &rightNum)

		leftNums = append(leftNums, leftNum)
		rightNums = append(rightNums, rightNum)

	}

	// Sort both arrays
	slices.Sort(leftNums)
	slices.Sort(rightNums)

	sum := 0

	for i, value := range leftNums {
		difference := value - rightNums[i]
		sum += int(math.Abs(float64(difference)))

	}

	return sum
}
