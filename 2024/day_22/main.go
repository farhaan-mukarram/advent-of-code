package main

import "fmt"

func main() {
	file := "input.txt"

	lines, err := parseInput(file)
	if err != nil {
		message := fmt.Sprintf("%v", err)
		panic(message)
	}

	fmt.Printf("Sum of each buyer's 2000th secret number = %d\n", part1(lines))
}
