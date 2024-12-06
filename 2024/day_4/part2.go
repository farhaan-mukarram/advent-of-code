package main

func searchDiagonals(rowStart int, colStart int, grid []string) int {
	seqCount := 0

	char := GetGridItem(rowStart, colStart, grid)

	if areStringsEqual(char, "A") {
		northWestChar := GetGridItem(rowStart-1, colStart-1, grid)
		southEastChar := GetGridItem(rowStart+1, colStart+1, grid)
		northEastChar := GetGridItem(rowStart-1, colStart+1, grid)
		southWestChar := GetGridItem(rowStart+1, colStart-1, grid)

		/*
		   M.S
		   .A.
		   M.S
		*/
		if areStringsEqual(northWestChar, "M") && areStringsEqual(southEastChar, "S") && areStringsEqual(northEastChar, "S") && areStringsEqual(southWestChar, "M") {
			seqCount += 1
		}

		/*
		   M.M
		   .A.
		   S.S
		*/
		if areStringsEqual(northWestChar, "M") && areStringsEqual(southEastChar, "S") && areStringsEqual(northEastChar, "M") && areStringsEqual(southWestChar, "S") {
			seqCount += 1
		}

		/*
		   S.S
		   .A.
		   M.M
		*/
		if areStringsEqual(northWestChar, "S") && areStringsEqual(southEastChar, "M") && areStringsEqual(northEastChar, "S") && areStringsEqual(southWestChar, "M") {
			seqCount += 1
		}

		/*
		   S.M
		   .A.
		   S.M
		*/
		if areStringsEqual(northWestChar, "S") && areStringsEqual(southEastChar, "M") && areStringsEqual(northEastChar, "M") && areStringsEqual(southWestChar, "S") {
			seqCount += 1
		}
	}

	return seqCount
}

func getCrossXmasCount(grid []string) int {
	totalCount := 0
	rowsLength := len(grid)
	colLength := len(grid)

	for i := 0; i < rowsLength; i++ {
		for j := 0; j < colLength; j++ {
			xmasCount := searchDiagonals(i, j, grid)
			totalCount += xmasCount
		}
	}

	return totalCount
}

func part2(lines []string) int {
	xmasCount := getCrossXmasCount(lines)
	return xmasCount
}
