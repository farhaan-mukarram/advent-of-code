package main

import "fmt"

func main() {
	file := "input.txt"

	lines, err := ParseInput(file)
	if err != nil {
		message := fmt.Sprintf("%v", err)
		panic(message)
	}

	fmt.Printf("Stone count after 25 blinks is %d\n", part1(lines[0]))
	fmt.Printf("Stone count after 75 blinks is %d\n", part2(lines[0]))
}
