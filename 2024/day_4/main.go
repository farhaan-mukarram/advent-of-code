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

	fmt.Printf("XMAS count is %d\n", part1(lines))
}
