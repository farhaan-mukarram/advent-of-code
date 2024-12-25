package main

import (
	"math"
	"os"
	"strconv"
	"strings"
)

func parseInput(filename string) ([]string, error) {
	contents, err := os.ReadFile(filename)

	var lines []string

	if err != nil {
		return lines, err
	}

	lines = strings.Split(strings.TrimSpace(string(contents)), "\n\n")

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

func isInt(x float64) bool {
	const epsilon = 1e-2 // Margin of error

	_, frac := math.Modf(math.Abs(x))

	return frac < epsilon || frac > 1.0-epsilon
}
