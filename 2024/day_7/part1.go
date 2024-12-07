package main

import (
	"regexp"
	"strings"
)

var operatorCombinations []string

func generateCombinations(currentString string, combinationLength int) {
	if len(currentString) == combinationLength {
		operatorCombinations = append(operatorCombinations, currentString)
	} else {
		generateCombinations(currentString+"+", combinationLength)
		generateCombinations(currentString+"*", combinationLength)
	}
}

func checkEquationPossibilities(combinations []string, lhs string, rhs string) bool {
	testValue := toInt(strings.TrimSpace(lhs))

	// extract numbers from the rhs
	re := regexp.MustCompile(`\d+`)
	nums := toIntArray(re.FindAllString(rhs, -1))

	// iterate over combinations, generate and evaluate equations
	for _, combination := range combinations {
		operatorIndex := 0

		res := 0
		for i := 0; i < len(nums)-1; i++ {
			isLastItem := i == len(nums)-2
			var operand_one, operand_two int
			if i == 0 {
				operand_one = nums[i]
			} else {
				operand_one = res
			}

			operand_two = nums[i+1]
			operator := string(combination[operatorIndex])
			operatorIndex++

			res = calculateResult(operand_one, operand_two, operator)

			// Last num processed, check total
			if isLastItem && res == testValue {
				return true
			}

		}
	}

	return false
}

func calculateResult(n1 int, n2 int, operator string) int {
	switch operator {
	case "+":
		return n1 + n2

	case "*":
		return n1 * n2

	}
	return 0
}

func part1(lines []string) int {
	calibrationSum := 0
	wrong_lines := 0

	for _, line := range lines {
		res := strings.Split(line, ":")
		result, operands := res[0], strings.TrimSpace(res[1])

		// determine number of operators by counting whitespace
		re := regexp.MustCompile(`\s`)
		numOfOperators := len(re.FindAllString(operands, -1))

		// clear the combinations array
		operatorCombinations = operatorCombinations[:0]

		// generator operator combinations
		generateCombinations("", numOfOperators)

		isEquationSolvable := checkEquationPossibilities(operatorCombinations, result, operands)

		if isEquationSolvable {
			calibrationSum += toInt(strings.TrimSpace(result))
		} else {
			wrong_lines++
		}

	}

	return calibrationSum
}
