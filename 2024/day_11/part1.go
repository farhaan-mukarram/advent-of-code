package main

import (
	"strings"
)

func part1(input string) int {
	var stoneStates []string
	blinkCount := 25
	stoneState := input

	for i := 0; i < blinkCount; i++ {
		stoneState = transformStones(stoneState)
		stoneStates = append(stoneStates, stoneState)
	}

	finalStoneState := stoneStates[len(stoneStates)-1]
	stonesArr := strings.Split(finalStoneState, " ")

	return len(stonesArr)
}
