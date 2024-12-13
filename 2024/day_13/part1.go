package main

import (
	"fmt"
	"strings"
)

const (
	A_BUTTON_PRESS_COST = 3
	B_BUTTON_PRESS_COST = 1
)

func part1(lines []string) int {
	var A0, A1, B0, B1, X0, X1, numOfTokens int

	for _, line := range lines {
		input := strings.Split(line, "\n")
		fmt.Sscanf(string(input[0]), "Button A: X+%d, Y+%d\n", &A0, &A1)
		fmt.Sscanf(string(input[1]), "Button B: X+%d, Y+%d\n", &B0, &B1)
		fmt.Sscanf(string(input[2]), "Prize: X=%d, Y=%d\n", &X0, &X1)

		for i := 0; i < 100; i++ {
			for j := 0; j < 100; j++ {
				R0 := (i * A0) + (j * B0)
				R1 := (i * A1) + (j * B1)

				if R0 == X0 && R1 == X1 {
					numOfTokens += (i * A_BUTTON_PRESS_COST) + (j * B_BUTTON_PRESS_COST)
				}
			}
		}

	}
	return numOfTokens
}
