package main

import (
	"fmt"
	"os"
	"strings"
)

func parseInput(filename string) ([]string, error) {
	contents, err := os.ReadFile(filename)

	var lines []string

	if err != nil {
		return lines, err
	}

	lines = strings.Split(strings.TrimSpace(string(contents)), "\n")

	return lines, nil
}

func main() {
	file := "input.txt"

	lines, err := ParseInput(file)
	if err != nil {
		message := fmt.Sprintf("%v", err)
		panic(message)
	}

	uniqueAntinodeLocationsPartOne := part1(lines)
	fmt.Printf("Unique antinode locations (part 1) = %d\n", uniqueAntinodeLocationsPartOne)
}
