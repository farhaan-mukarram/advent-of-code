package main

import (
	"strconv"
	"strings"
)

func part1(diskMap string) int {
	idCount := 0
	var blockArray []string

	for i, char := range diskMap {
		ch := string(char)
		num := toInt(ch)

		if isEven(i) {
			for i := 0; i < num; i++ {
				blockArray = append(blockArray, strconv.Itoa(idCount))
			}
			idCount++
		} else {
			for i := 0; i < num; i++ {
				blockArray = append(blockArray, ".")
			}
		}
	}

	for !isFreeSpaceContiguous(strings.Join(blockArray, "")) {
		freeSpaceIndex := 0
		numIndex := 0

		// Find index of free space
		for i, c := range blockArray {
			ch := string(c)

			if areStringsEqual(ch, ".") {
				freeSpaceIndex = i
				break
			}

		}

		// Find index of number from right side
		for i := len(blockArray) - 1; i >= 0; i-- {
			ch := blockArray[i]

			if !areStringsEqual(ch, ".") {
				numIndex = i
				break
			}

		}

		blockArray[freeSpaceIndex], blockArray[numIndex] = blockArray[numIndex], blockArray[freeSpaceIndex]

	}

	checksum := getFileChecksum(blockArray)

	return checksum
}
