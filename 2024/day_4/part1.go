package main

const (
	SEQUENCE = "XMAS"
)

func searchGridForWord(rowStart int, colStart int, grid []string) int {
	seqCount := 0
	seqLength := len(SEQUENCE)

	// search north
	seq := ""
	for i := 0; i < seqLength; i++ {
		char := GetGridItem(rowStart-i, colStart, grid)
		seq += string(char)
	}

	if areStringsEqual(seq, SEQUENCE) {
		seqCount += 1
	}

	// search north-east
	seq = ""
	for i := 0; i < seqLength; i++ {
		char := GetGridItem(rowStart-i, colStart+i, grid)
		seq += string(char)
	}

	if areStringsEqual(seq, SEQUENCE) {
		seqCount += 1
	}

	// search east
	seq = ""
	for i := 0; i < seqLength; i++ {
		char := GetGridItem(rowStart, colStart+i, grid)
		seq += string(char)
	}

	if areStringsEqual(seq, SEQUENCE) {
		seqCount += 1
	}

	// search south-east
	seq = ""
	for i := 0; i < seqLength; i++ {
		char := GetGridItem(rowStart+i, colStart+i, grid)
		seq += string(char)
	}

	if areStringsEqual(seq, SEQUENCE) {
		seqCount += 1
	}

	// search south
	seq = ""
	for i := 0; i < seqLength; i++ {
		char := GetGridItem(rowStart+i, colStart, grid)
		seq += string(char)
	}

	if areStringsEqual(seq, SEQUENCE) {
		seqCount += 1
	}

	// search south-west
	seq = ""
	for i := 0; i < seqLength; i++ {
		char := GetGridItem(rowStart+i, colStart-i, grid)
		seq += string(char)
	}

	if areStringsEqual(seq, SEQUENCE) {
		seqCount += 1
	}

	// search west
	seq = ""
	for i := 0; i < seqLength; i++ {
		char := GetGridItem(rowStart, colStart-i, grid)
		seq += string(char)
	}

	if areStringsEqual(seq, SEQUENCE) {
		seqCount += 1
	}

	// search north-west
	seq = ""
	for i := 0; i < seqLength; i++ {
		char := GetGridItem(rowStart-i, colStart-i, grid)
		seq += string(char)
	}

	if areStringsEqual(seq, SEQUENCE) {
		seqCount += 1
	}

	return seqCount
}

func getXmasCount(grid []string) int {
	totalCount := 0
	rowsLength := len(grid)
	colLength := len(grid)

	for i := 0; i < rowsLength; i++ {
		for j := 0; j < colLength; j++ {
			xmasCount := searchGridForWord(i, j, grid)
			totalCount += xmasCount
		}
	}

	return totalCount
}

func part1(lines []string) int {
	xmasCount := getXmasCount(lines)
	return xmasCount
}
