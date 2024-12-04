package main

import (
	"fmt"
	"regexp"
)

func part1(input string) int {
	sum_of_prods := 0
	var num_1, num_2 int

	re := regexp.MustCompile(MUL_INSTRUCTION_REGEX_PATTERN)
	instructions := re.FindAllString(input, -1)

	for _, instruction := range instructions {
		fmt.Sscanf(instruction, MUL_INSTRUCTION_SCAN_PATTERN, &num_1, &num_2)
		sum_of_prods += num_1 * num_2
	}

	return sum_of_prods
}
