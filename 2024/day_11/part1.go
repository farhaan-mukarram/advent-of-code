package main

import (
	"fmt"
	"strings"
)

func transformStones(input string) string {
	stonesArr := strings.Split(input, " ")
	var newStonesArr []string

	for _, stone := range stonesArr {
		numOfDigits := len(stone)
		num := toInt(stone)

		if num == 0 {
			newStonesArr = append(newStonesArr, "1")
		} else if isEven(numOfDigits) {
			// Split into left and right parts
			leftPart, rightPart := stone[:numOfDigits/2], stone[numOfDigits/2:]

			// Convert to numbers to remove leading zeroes
			leftNum, rightNum := toInt(leftPart), toInt(rightPart)

			// TODO: Figure out a better way for this
			newStonesArr = append(newStonesArr, toString(leftNum))
			newStonesArr = append(newStonesArr, toString(rightNum))
		} else {
			num *= 2024
			newStonesArr = append(newStonesArr, toString(num))
		}
	}

	return strings.Join(newStonesArr, " ")
}

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
	fmt.Printf("Final stone state length %d\n", len(stonesArr))

	return 0
}
