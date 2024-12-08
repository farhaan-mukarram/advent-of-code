package main

import (
	"unicode"
)

func part2(lines []string) int {
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

			rise := coordOne.row - coordTwo.row
			run := coordOne.col - coordTwo.col

			var antinodeLocationOne, antinodeLocationTwo Coordinate

			antinodeLocationOne = Coordinate{row: coordOne.row, col: coordOne.col}
			antinodeLocationTwo = Coordinate{row: coordTwo.row, col: coordTwo.col}

			for liesWithinGrid(antinodeLocationOne, lines) {
				if !(doesListContainCoordinate(antinodeLocationOne, antinodeLocations)) {
					antinodeLocations = append(antinodeLocations, antinodeLocationOne)
				}

				antinodeLocationOne = Coordinate{row: antinodeLocationOne.row + rise, col: antinodeLocationOne.col + run}
			}

			for liesWithinGrid(antinodeLocationTwo, lines) {
				if !(doesListContainCoordinate(antinodeLocationTwo, antinodeLocations)) {
					antinodeLocations = append(antinodeLocations, antinodeLocationTwo)
				}

				antinodeLocationTwo = Coordinate{row: antinodeLocationTwo.row - rise, col: antinodeLocationTwo.col - run}
			}

		}

	}

	return len(antinodeLocations)
}
