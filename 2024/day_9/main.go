package main

import "fmt"

func main() {
	file := "input.txt"

	lines, err := ParseInput(file)
	if err != nil {
		message := fmt.Sprintf("%v", err)
		panic(message)
	}

	input := lines[0]

	fileCheckSumPartOne := part1(input)
	fileCheckSumPartTwo := part2(input)

	fmt.Printf("File checksum (part 1) = %d", fileCheckSumPartOne)
	fmt.Printf("File checksum (part 2) = %d", fileCheckSumPartTwo)
}
