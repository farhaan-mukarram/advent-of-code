package main

import (
	"os"
	"strings"
)

func ParseInput(filename string) (string, error) {
	contents, err := os.ReadFile(filename)

	var data string

	if err != nil {
		return data, err
	}

	data = strings.TrimSpace(string(contents))

	return data, nil
}
