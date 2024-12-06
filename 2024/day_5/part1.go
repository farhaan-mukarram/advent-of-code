package main

import (
	"fmt"
	"strings"
)

func isUpdateCorrectlyOrdered(nums []int, rulesString string) bool {
	numsLength := len(nums)
	isCorrectlyOrdered := true

	for i := 0; i < numsLength; i++ {
		for j := i + 1; j < numsLength; j++ {
			// fmt.Printf("\t Comparing %d and %d\n", nums[i], nums[j])
			rule := fmt.Sprintf("%d|%d", nums[i], nums[j])

			isCorrectlyOrdered = strings.Contains(rulesString, rule)

			if !isCorrectlyOrdered {
				return false
			}
		}
	}

	return isCorrectlyOrdered
}

func part1(lines []string) int {
	medianSum := 0
	rulesString, updatesString := lines[0], lines[1]
	updates := strings.Split(updatesString, "\n")

	for _, update := range updates {
		nums := toIntArray(strings.Split(update, ","))
		isCorrectlyOrdered := isUpdateCorrectlyOrdered(nums, rulesString)
		// fmt.Printf("Update = %s, correctly ordered = %v\n", update, isCorrectlyOrdered)

		if isCorrectlyOrdered {
			median := nums[len(nums)/2]
			medianSum += median
		}
	}

	return medianSum
}
