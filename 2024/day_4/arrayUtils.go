package main

import "strings"

func GetGridItem(rowIdx int, colIdx int, grid []string) string {
	numOfRows := len(grid)
	numOfColumns := len(grid[0])

	if rowIdx < 0 || rowIdx > numOfRows-1 {
		return "."
	}

	if colIdx < 0 || colIdx > numOfColumns-1 {
		return "."
	}

	char := grid[rowIdx][colIdx]

	return string(char)
}

func areStringsEqual(s1 string, s2 string) bool {
	return strings.Compare(s1, s2) == 0
}
