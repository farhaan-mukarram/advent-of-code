package main

import "fmt"

func main() {
	file := "input.txt"

	input, err := ParseInput(file)
	if err != nil {
		message := fmt.Sprintf("%v", err)
		panic(message)
	}

	allMulResult := part1(input)
	enabledMulResult := part2(input)

	fmt.Printf("Result of adding all multiplications = %d\n", allMulResult)
	fmt.Printf("Result of adding only enabled multiplications = %d\n", enabledMulResult)
}
