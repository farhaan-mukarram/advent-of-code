package main

import "fmt"

func main() {
	file := "input.txt"

	lines, err := parseInput(file)
	if err != nil {
		message := fmt.Sprintf("%v", err)
		panic(message)
	}

	fmt.Printf("The final output string (part 1) is = %s\n", part1(lines))
}
