package main

import (
	"regexp"
)

func isSeqSafe(seq []int) bool {
	isIncreasing := isSequenceIncreasing(seq)
	isDecreasing := isSequenceDecreasing(seq)

	if !isDecreasing && !isIncreasing {
		return false
	}

	isSafe := true

	for i := 1; i < len(seq); i++ {
		curr := seq[i]
		prev := seq[i-1]
		delta := abs(curr - prev)

		if delta < MIN || delta > MAX {
			isSafe = false
			break
		}
	}

	return isSafe
}

func isSequenceSafeWithBadLevel(seq []int) bool {
	isSafe := false

	for i := 0; i < len(seq); i++ {
		// remove the ith element and check the remaining sequence
		tempArr := remove(seq, i)

		if isSeqSafe(tempArr) {
			isSafe = true
			break
		}

	}

	return isSafe
}

func part2(lines []string) int {
	numOfSafeReports := 0

	re := regexp.MustCompile(`\d+`)

	for _, line := range lines {
		nums := toIntArray(re.FindAllString(line, -1))

		isSeqSafe := isSequenceSafeWithBadLevel(nums)

		if isSeqSafe {
			numOfSafeReports++
		}

	}

	return numOfSafeReports
}
