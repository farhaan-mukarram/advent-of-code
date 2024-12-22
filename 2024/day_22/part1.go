package main

func part1(lines []string) int {
	sumOfTwoThousandnthSecretNumber := 0

	for _, line := range lines {
		secretNumber := toInt(line)

		for i := 0; i < 2000; i++ {
			secretNumber = computeSecretNumber(secretNumber)
		}

		sumOfTwoThousandnthSecretNumber += secretNumber
	}

	return sumOfTwoThousandnthSecretNumber
}
