package main

import (
	"regexp"
)

func isSequenceIncreasing(seq []int) bool {
	isIncreasing := true

	for i := 1; i < len(seq); i++ {

		curr := seq[i]
		prev := seq[i-1]

		if curr <= prev {
			isIncreasing = false
		}
	}

	return isIncreasing
}

func isSequenceDecreasing(seq []int) bool {
	isDecreasing := true

	for i := 1; i < len(seq); i++ {

		curr := seq[i]
		prev := seq[i-1]

		if curr >= prev {
			isDecreasing = false
		}
	}

	return isDecreasing
}

func isSequenceSafe(seq []int) bool {
	isSafe := true
	isIncreasing := isSequenceIncreasing(seq)
	isDecreasing := isSequenceDecreasing(seq)

	if isIncreasing || isDecreasing {
		for i := 1; i < len(seq); i++ {

			curr := seq[i]
			prev := seq[i-1]
			delta := abs(curr - prev)

			if delta < MIN || delta > MAX {
				isSafe = false
			}
		}
	} else {
		isSafe = false
	}

	return isSafe
}

func part1(lines []string) int {
	numOfSafeReports := 0

	re := regexp.MustCompile(`\d+`)

	for _, line := range lines {
		nums := toIntArray(re.FindAllString(line, -1))

		isSeqSafe := isSequenceSafe(nums)

		if isSeqSafe {
			numOfSafeReports++
		}

	}

	return numOfSafeReports
}
