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

	fmt.Printf("Sum of middle paged numbers of correctly ordered updates is = %d", part1(lines))
}
