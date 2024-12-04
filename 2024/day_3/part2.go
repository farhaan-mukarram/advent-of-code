package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func parseInput(filename string) (string, error) {
	contents, err := os.ReadFile(filename)

	var lines string

	if err != nil {
		return lines, err
	}

	lines = strings.TrimSpace(string(contents))

	return lines, nil
}

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

func part2(filename string) int {
	sum_of_prods := 0

	input, _ := parseInput(filename)

	isMulDisabled := false
	buffer := ""

	for _, char := range input {
		if !isValidChar(char) {
			buffer = ""
		} else {
			buffer += string(char)

			isMulInstruction, _ := regexp.MatchString(`(mul\(\d+,\d+\))`, buffer)
			isDoInstruction, _ := regexp.MatchString(`do\(\)`, buffer)
			isDontInstruction, _ := regexp.MatchString(`don't\(\)`, buffer)

			if isMulInstruction && !isMulDisabled {
				var num_1, num_2 int

				re := regexp.MustCompile(`(mul\(\d+,\d+\))`)
				res := re.FindAllStringSubmatch(buffer, -1)
				matchedPart := res[0][1]
				fmt.Sscanf(matchedPart, "mul(%d,%d)", &num_1, &num_2)

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
