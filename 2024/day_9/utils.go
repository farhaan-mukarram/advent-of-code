package main

import (
	"regexp"
	"strconv"
	"strings"
)

func isEven(num int) bool {
	return num%2 == 0
}

func toInt(s string) int {
	num, _ := strconv.Atoi(s)

	return num
}

func areStringsEqual(s1 string, s2 string) bool {
	return strings.Compare(s1, s2) == 0
}

func isFreeSpaceContiguous(s string) bool {
	re := regexp.MustCompile(`^(\d+)(\.+)$`)
	return re.MatchString(s)
}

func getFileChecksum(block []string) int {
	checksum := 0

	for i, c := range block {
		ch := string(c)

		if areStringsEqual(ch, ".") {
			continue
		}

		fileIdNumber := toInt(ch)
		checksum += i * fileIdNumber
	}

	return checksum
}