package main

import (
	"reflect"
	"unicode"
)

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

func part1(lines []string) int {
	var antinodeLocations []Coordinate

	antennaMap := make(map[string][]Coordinate)

	for i, line := range lines {
		for j, char := range line {
			// Check possible antenna
			if unicode.IsDigit(char) || unicode.IsLetter(char) {
				antennaCoords := Coordinate{row: i, col: j}
				var coordsList []Coordinate

				// Check map
				existingCoordsList, isPresent := antennaMap[string(char)]

				// If coords already exist, use existing list
				if isPresent {
					coordsList = existingCoordsList
				}

				// Update antennaMap
				coordsList = append(coordsList, antennaCoords)
				antennaMap[string(char)] = coordsList

			}
		}
	}

	for _, coords := range antennaMap {

		coordsToProcess := getCoordinatesToProcess(coords)

		for _, coordPair := range coordsToProcess {
			coordOne, coordTwo := coordPair[0], coordPair[1]
			// fmt.Printf("Coordinates = %v\n", coordPair)

			rise := coordOne.row - coordTwo.row
			run := coordOne.col - coordTwo.col

			var antinodeLocationOne, antinodeLocationTwo Coordinate
			var c1, c2 Coordinate

			c1 = Coordinate{row: coordOne.row + rise, col: coordOne.col + run}
			c2 = Coordinate{row: coordTwo.row - rise, col: coordTwo.col - run}

			antinodeLocationOne = c1
			antinodeLocationTwo = c2
			// fmt.Printf("\tantinode 1 = %d, antinode 2 = %d\n", c1, c2)

			if liesWithinGrid(antinodeLocationOne, lines) && !doesListContainCoordinate(antinodeLocationOne, antinodeLocations) {
				antinodeLocations = append(antinodeLocations, antinodeLocationOne)
			}

			if liesWithinGrid(antinodeLocationTwo, lines) && !doesListContainCoordinate(antinodeLocationTwo, antinodeLocations) {
				antinodeLocations = append(antinodeLocations, antinodeLocationTwo)
			}

			// fmt.Printf("\tantinode locations = %v\n\n", antinodeLocations)
		}

	}

	return len(antinodeLocations)
}
