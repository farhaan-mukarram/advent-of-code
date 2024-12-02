package main

import (
	"fmt"
)

func main() {
	file := "input.txt"

	lines, err := ParseInput(file)
	if err != nil {
		message := fmt.Sprintf("%v", err)
		panic(message)
	}

	numOfSafeReports := part1(lines)
	numOfSafeReportsWithOneBadLevel := part2(lines)

	fmt.Printf("The number of safe reports is : %d\n", numOfSafeReports)
	fmt.Printf("The number of safe reports is (with bad levels) : %d\n", numOfSafeReportsWithOneBadLevel)
}
