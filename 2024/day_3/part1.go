package main

import (
	"fmt"
	"regexp"
)

func compute_row_sum(row []string) int {
	row_sum := 0

	for _, instruction := range row {
		var num_1, num_2 int
		fmt.Sscanf(instruction, "mul(%d,%d)", &num_1, &num_2)

		row_sum += num_1 * num_2

	}

	return row_sum
}

func part1(lines []string) int {
	sum_of_prods := 0

	for _, line := range lines {
		re := regexp.MustCompile(REGEX_PATTERN)
		instructions := re.FindAllString(line, -1)

		row_sum := compute_row_sum(instructions)
		sum_of_prods += row_sum

	}

	return sum_of_prods
}
