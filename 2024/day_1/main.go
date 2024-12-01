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

	distanceBetweenLists := part1(lines)
	similarityScore := part2(lines)

	fmt.Printf("Distance between lists = %d\n", distanceBetweenLists)
	fmt.Printf("Similarity score = %d\n", similarityScore)
}
