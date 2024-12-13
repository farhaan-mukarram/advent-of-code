package main

import (
	"os"
	"strings"
)

func parseInput(filename string) ([]string, error) {
	contents, err := os.ReadFile(filename)

	var lines []string

	if err != nil {
		return lines, err
	}

	lines = strings.Split(strings.TrimSpace(string(contents)), "\n\n")

	return lines, nil
}
