package main

import (
	"fmt"
	"strings"
)

// enum to store all opcodes
const (
	ADV = iota
	BXL
	BST
	JNZ
	BXC
	OUT
	BDV
	CDV
)

func part1(lines []string) string {
	var regA, regB, regC, outputString string
	var instructions []int
	var output []string
	instructionPtr := 0

	regInput := lines[0]
	programInput := lines[1]

	// Register details
	regDetails := strings.Split(regInput, "\n")
	fmt.Sscanf(regDetails[0], "Register A: %s", &regA)
	fmt.Sscanf(regDetails[1], "Register B: %s", &regB)
	fmt.Sscanf(regDetails[2], "Register C: %s", &regC)

	// Instruction details
	fmt.Sscanf(programInput, "Program: %s", &programInput)
	instructions = toIntArray(strings.Split(programInput, ","))

	for instructionPtr < len(instructions) {
		opcode := instructions[instructionPtr]

		switch opcode {
		case ADV:
			{
				var denominator, result, power int
				operand := instructions[instructionPtr+1]
				numerator := toInt(regA)

				// Combo operand
				if operand >= 0 && operand <= 3 {
					power = operand
				} else { // 4 - 7
					switch operand {
					// A
					case 4:
						power = toInt(regA)

					// B
					case 5:
						power = toInt(regB)

					// C
					case 6:
						power = toInt(regC)
					}
				}

				denominator = pow(2, power)

				// Compute result and save in regA
				result = numerator / denominator
				regA = toString(result)

				instructionPtr += 2

			}

		case BXL:
			{
				// Compute bitwise xor of regB and operand
				intB := toInt(regB)
				operand := instructions[instructionPtr+1]
				result := intB ^ operand

				// Update regB
				regB = toString(result)

				instructionPtr += 2

			}

		case BST:
			{
				operand := instructions[instructionPtr+1]
				var num, result int

				// Combo operand
				if operand >= 0 && operand <= 3 {
					num = operand
				} else { // 4 - 7
					switch operand {
					// A
					case 4:
						num = toInt(regA)

					// B
					case 5:
						num = toInt(regB)

					// C
					case 6:
						num = toInt(regC)
					}
				}

				// Compute modulo 8
				result = num % 8

				// Update regB
				regB = toString(result)

				instructionPtr += 2
			}

		case JNZ:
			{
				intA := toInt(regA)

				// Update ptr if regA != 0
				if intA != 0 {
					address := instructions[instructionPtr+1]

					instructionPtr = address
				} else {
					instructionPtr += 2
				}
			}

		case BXC:
			{
				// bitwise xor of registers b and c
				result := toInt(regB) ^ toInt(regC)
				regB = toString(result)

				instructionPtr += 2
			}

		case OUT:
			{

				operand := instructions[instructionPtr+1]
				var num, result int

				// Combo operand
				if operand >= 0 && operand <= 3 {
					num = operand
				} else { // 4 - 7
					switch operand {
					// A
					case 4:
						num = toInt(regA)

					// B
					case 5:
						num = toInt(regB)

					// C
					case 6:
						num = toInt(regC)
					}
				}

				// Compute modulo 8 and save
				result = num % 8
				output = append(output, toString(result))

				instructionPtr += 2
			}

		case BDV:
			{
				var denominator, result, power int
				operand := instructions[instructionPtr+1]
				numerator := toInt(regA)

				// Combo operand
				if operand >= 0 && operand <= 3 {
					power = operand
				} else { // 4 - 7
					switch operand {
					// A
					case 4:
						power = toInt(regA)

					// B
					case 5:
						power = toInt(regB)

					// C
					case 6:
						power = toInt(regC)
					}
				}

				denominator = pow(2, power)

				// Compute result and save in regB
				result = numerator / denominator
				regB = toString(result)

				instructionPtr += 2
			}

		case CDV:
			{
				var denominator, result, power int
				operand := instructions[instructionPtr+1]
				numerator := toInt(regA)

				// Combo operand
				if operand >= 0 && operand <= 3 {
					power = operand
				} else { // 4 - 7
					switch operand {
					// A
					case 4:
						power = toInt(regA)

					// B
					case 5:
						power = toInt(regB)

					// C
					case 6:
						power = toInt(regC)
					}
				}

				denominator = pow(2, power)

				// Compute result and save in regC
				result = numerator / denominator
				regC = toString(result)

				instructionPtr += 2
			}
		}
	}

	outputString = strings.Join(output, ",")

	return outputString
}
