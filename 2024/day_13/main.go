package main

import "fmt"

func main() {
	file := "input.txt"

	lines, err := parseInput(file)
	if err != nil {
		message := fmt.Sprintf("%v", err)
		panic(message)
	}

	fmt.Printf("Fewest tokens to win all possible prizes (part 1) = %d", part1(lines))
}
