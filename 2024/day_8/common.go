package main

import "reflect"

type Coordinate struct {
	row int
	col int
}

func doesListContainCoordinate(val Coordinate, arr []Coordinate) bool {
	containsCoordinate := false

	for _, c := range arr {
		if reflect.DeepEqual(c, val) {
			return true
		}
	}

	return containsCoordinate
}

func liesWithinGrid(val Coordinate, lines []string) bool {
	rowCount := len(lines)
	colCount := len(lines[0])

	if val.row < 0 || val.row > rowCount-1 {
		return false
	}

	if val.col < 0 || val.col > colCount-1 {
		return false
	}

	return true
}

func getCoordinatesToProcess(coords []Coordinate) [][]Coordinate {
	var processedCoords [][]Coordinate

	for i, coordOne := range coords {
		for j, coordTwo := range coords {

			if i == j {
				continue
			}

			shouldAppend := true

			// search if processed coords contain combination (ignore order)
			for _, c := range processedCoords {
				c1, c2 := c[0], c[1]

				if reflect.DeepEqual(c1, coordOne) && reflect.DeepEqual(c2, coordTwo) {
					shouldAppend = false
				} else if reflect.DeepEqual(c1, coordTwo) && reflect.DeepEqual(c2, coordOne) {
					shouldAppend = false
				}

			}

			// only append if combination not processed
			if shouldAppend {
				var coordsToAppend []Coordinate

				coordsToAppend = append(coordsToAppend, coordOne)
				coordsToAppend = append(coordsToAppend, coordTwo)

				processedCoords = append(processedCoords, coordsToAppend)
			}

		}
	}

	return processedCoords
}
