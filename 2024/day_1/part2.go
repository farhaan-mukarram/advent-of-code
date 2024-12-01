package main

import (
	"fmt"
)

func part2(lines []string) int {
	var leftNums, rightNums []int
	var leftNum, rightNum int

	// Map containing counts
	counts := make(map[int]int)

	// Populate left and right arrays with numbers
	for _, line := range lines {
		fmt.Sscanf(line, "%d %d", &leftNum, &rightNum)

		leftNums = append(leftNums, leftNum)
		rightNums = append(rightNums, rightNum)

		counts[rightNum]++
	}

	similarityScore := 0
	for _, num := range leftNums {
		similarityScore += num * counts[num]
	}

	return similarityScore
}
