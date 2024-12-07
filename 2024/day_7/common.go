package main

import (
	"regexp"
	"strings"
)

func generateCombinations(currentString string, combinationLength int, allowedOperators []string) {
	if len(currentString) == combinationLength {
		operatorCombinations = append(operatorCombinations, currentString)
	} else {
		for _, operator := range allowedOperators {
			generateCombinations(currentString+operator, combinationLength, allowedOperators)
		}
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
			var operandOne, operandTwo int

			// First operand is the first number at the start, prev result otherwise
			if i == 0 {
				operandOne = nums[i]
			} else {
				operandOne = res
			}

			operandTwo = nums[i+1]

			// Determining the operator to use
			operator := string(combination[operatorIndex])
			operatorIndex++

			res = calculateResult(operandOne, operandTwo, operator)
		}

		// Last num processed, check total
		if res == testValue {
			return true
		}
	}

	return false
}

func calculateResult(n1 int, n2 int, operator string) int {
	switch operator {
	case ADD_OPERATOR:
		return n1 + n2

	case MULTIPLY_OPERATOR:
		return n1 * n2

	case CONCAT_OPERATOR:
		numStrOne := toString(n1)
		numStrTwo := toString(n2)
		combinedNum := numStrOne + numStrTwo

		return toInt(combinedNum)
	}

	return 0
}

func determineCalibrationSum(lines []string, operators []string) int {
	calibrationSum := 0

	for _, line := range lines {
		res := strings.Split(line, ":")
		result, operands := res[0], strings.TrimSpace(res[1])

		// determine number of operators by counting whitespace
		re := regexp.MustCompile(`\s`)
		numOfOperators := len(re.FindAllString(operands, -1))

		// clear the combinations array
		operatorCombinations = operatorCombinations[:0]

		// generator operator combinations
		generateCombinations("", numOfOperators, operators)

		isEquationSolvable := checkEquationPossibilities(operatorCombinations, result, operands)

		if isEquationSolvable {
			calibrationSum += toInt(strings.TrimSpace(result))
		}
	}

	return calibrationSum
}
