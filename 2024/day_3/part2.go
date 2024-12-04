package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func isValidChar(r rune) bool {
	if unicode.IsDigit(r) {
		return true
	}

	char := string(r)
	if strings.Compare("d", char) == 0 {
		return true
	}

	if strings.Compare("o", char) == 0 {
		return true
	}

	if strings.Compare("n", char) == 0 {
		return true
	}

	if strings.Compare("t", char) == 0 {
		return true
	}

	if strings.Compare("m", char) == 0 {
		return true
	}

	if strings.Compare("u", char) == 0 {
		return true
	}

	if strings.Compare("l", char) == 0 {
		return true
	}

	if strings.Compare("(", char) == 0 {
		return true
	}

	if strings.Compare(")", char) == 0 {
		return true
	}

	if strings.Compare(",", char) == 0 {
		return true
	}

	if strings.Compare("'", char) == 0 {
		return true
	}

	return false
}

func part2(input string) int {
	sum_of_prods := 0

	isMulDisabled := false
	buffer := ""

	for _, char := range input {
		if !isValidChar(char) {
			buffer = ""
		} else {
			buffer += string(char)

			isMulInstruction, _ := regexp.MatchString(MUL_INSTRUCTION_REGEX_PATTERN, buffer)
			isDoInstruction, _ := regexp.MatchString(`do\(\)`, buffer)
			isDontInstruction, _ := regexp.MatchString(`don't\(\)`, buffer)

			if isMulInstruction && !isMulDisabled {
				var num_1, num_2 int

				re := regexp.MustCompile(MUL_INSTRUCTION_REGEX_PATTERN)
				res := re.FindAllStringSubmatch(buffer, -1)

				matchedPart := res[0][1]
				fmt.Sscanf(matchedPart, MUL_INSTRUCTION_SCAN_PATTERN, &num_1, &num_2)

				sum_of_prods += num_1 * num_2
				buffer = ""
			} else if isDoInstruction {
				isMulDisabled = false
				buffer = ""
			} else if isDontInstruction {
				isMulDisabled = true
				buffer = ""
			}
		}
	}

	return sum_of_prods
}
