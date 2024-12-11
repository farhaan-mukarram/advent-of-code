package main

import (
	"strings"
)

func updateStonesMap(stonesMap map[int]int) map[int]int {
	newStonesMap := make(map[int]int)
	for stoneValue, stoneCount := range stonesMap {
		stoneValueString := toString(stoneValue)
		numOfDigits := len(stoneValueString)

		if stoneValue == 0 {
			newStonesMap[1] = newStonesMap[1] + stoneCount
		} else if isEven(numOfDigits) {

			// Split into left and right parts
			leftPart, rightPart := stoneValueString[:numOfDigits/2], stoneValueString[numOfDigits/2:]

			// Convert to numbers to remove leading zeroes
			leftNum, rightNum := toInt(leftPart), toInt(rightPart)

			newStonesMap[leftNum] = newStonesMap[leftNum] + stoneCount
			newStonesMap[rightNum] = newStonesMap[rightNum] + stoneCount

		} else {
			stoneValue *= 2024
			newStonesMap[stoneValue] = newStonesMap[stoneValue] + stoneCount
		}

	}

	return newStonesMap
}

func initStonesMap(input string) map[int]int {
	stonesMap := make(map[int]int)
	stonesArr := strings.Split(input, " ")
	for _, stone := range stonesArr {
		stoneValue := toInt(stone)

		stonesMap[stoneValue]++

	}

	return stonesMap
}

func part2(input string) int {
	blinkCount := 75

	stonesMap := initStonesMap(input)

	for i := 0; i < blinkCount; i++ {
		stonesMap = updateStonesMap(stonesMap)
	}

	totalStoneCount := 0

	for _, stoneCount := range stonesMap {
		totalStoneCount += stoneCount
	}

	return totalStoneCount
}
