package main

import (
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
	num, _ := strconv.ParseInt(s, 10, 0)

	return int(num)
}

func toString(n int) string {
	return strconv.Itoa(n)
}

func mix(n1 int, n2 int) int {
	// bitwise xor
	return n1 ^ n2
}

func prune(n int) int {
	// remainder
	return n % 16777216
}

func computeSecretNumber(secretNumber int) int {
	nextSecretNumber := secretNumber
	var a, b int

	// step 1: multiply by 64, mix with original and prune
	a = nextSecretNumber * 64
	b = mix(nextSecretNumber, a)
	nextSecretNumber = prune(b)

	// step 2: divide by 32, mix with original and prune
	a = nextSecretNumber / 32
	b = mix(nextSecretNumber, a)
	nextSecretNumber = prune(b)

	// step 3: multiply by 2048, mix with original and prune
	a = nextSecretNumber * 2048
	b = mix(nextSecretNumber, a)
	nextSecretNumber = prune(b)

	return nextSecretNumber
}
