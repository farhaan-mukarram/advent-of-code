package main

import "fmt"

func main() {
	file := "input.txt"

	lines, err := parseInput(file)
	if err != nil {
		message := fmt.Sprintf("%v", err)
		panic(message)
	}

	fmt.Printf("Safety factor (part 1) = %d\n", part1(lines))
}
