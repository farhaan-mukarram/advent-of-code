package main

import (
	"os"
	"strconv"
	"strings"
)

func ParseInput(filename string) ([]string, error) {
	contents, err := os.ReadFile(filename)

	var lines []string

	if err != nil {
		return lines, err
	}

	lines = strings.Split(strings.TrimSpace(string(contents)), "\n")

	return lines, nil
}

func toIntArray(arr []string) []int {
	var res []int

	for _, val := range arr {
		num := toInt(val)
		res = append(res, num)

	}

	return res
}

func toInt(s string) int {
	num, _ := strconv.Atoi(s)

	return num
}

func toString(n int) string {
	return strconv.Itoa(n)
}

func isEven(n int) bool {
	return n%2 == 0
}

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
