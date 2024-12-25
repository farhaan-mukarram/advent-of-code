package main

import (
	"fmt"
	"math"
	"strings"

	"gonum.org/v1/gonum/mat"
)

func part2(lines []string) int {
	var A0, A1, B0, B1, X0, X1, numOfTokens int

	for _, line := range lines {
		input := strings.Split(line, "\n")
		fmt.Sscanf(string(input[0]), "Button A: X+%d, Y+%d\n", &A0, &A1)
		fmt.Sscanf(string(input[1]), "Button B: X+%d, Y+%d\n", &B0, &B1)
		fmt.Sscanf(string(input[2]), "Prize: X=%d, Y=%d\n", &X0, &X1)

		X0 += OFFSET
		X1 += OFFSET

		coeffMatrix := mat.NewDense(2, 2, []float64{float64(A0), float64(B0), float64(A1), float64(B1)})

		matA := mat.DenseCopyOf(coeffMatrix)
		matA.SetCol(0, []float64{float64(X0), float64(X1)})

		matB := mat.DenseCopyOf(coeffMatrix)
		matB.SetCol(1, []float64{float64(X0), float64(X1)})

		D := mat.Det(coeffMatrix)
		Dx := mat.Det(matA)
		Dy := mat.Det(matB)

		var x, y float64

		if D != 0 {
			x = Dx / D
			y = Dy / D

			if isInt(x) && isInt(y) {
				roundedX := int(math.Round(x))
				roundedY := int(math.Round(y))

				numOfTokens += (roundedX * A_BUTTON_PRESS_COST) + (roundedY * B_BUTTON_PRESS_COST)
			}
		}

	}

	return numOfTokens
}
