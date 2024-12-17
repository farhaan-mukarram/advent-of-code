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
	num, _ := strconv.ParseInt(s, 10, 0)

	return int(num)
}

func toString(n int) string {
	return strconv.Itoa(n)
}

func pow(base int, power int) int {
	return int(math.Pow(float64(base), float64(power)))
}
