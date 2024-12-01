package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func parseInput(filename string) ([]string, error) {
	contents, err := os.ReadFile(filename)

	var lines []string

	if err != nil {
		return lines, err
	}

	lines = strings.Split(strings.TrimSpace(string(contents)), "\n")

	return lines, nil
}

func main() {
	file := "input.txt"

	lines, err := parseInput(file)
	if err != nil {
		message := fmt.Sprintf("%v", err)
		panic(message)
	}

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

	fmt.Println(sum)
}
