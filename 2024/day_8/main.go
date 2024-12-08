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

	uniqueAntinodeLocationsPartOne := part1(lines)
	uniqueAntinodeLocationsPartTwo := part2(lines)

	fmt.Printf("Unique antinode locations (part 1) = %d\n", uniqueAntinodeLocationsPartOne)
	fmt.Printf("Unique antinode locations (part 2) = %d\n", uniqueAntinodeLocationsPartTwo)
}
