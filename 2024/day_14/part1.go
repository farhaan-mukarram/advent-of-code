package main

import (
	"regexp"
)

func updateRobotLocation(robotInfo RobotInfo, rowMax int, colMax int) RobotInfo {
	px := robotInfo.Position.x
	py := robotInfo.Position.y
	vx := robotInfo.Velocity.x
	vy := robotInfo.Velocity.y

	pxUpdated := px + vx
	pyUpdated := py + vy

	// x-axis/column teleport logic
	if pxUpdated > colMax-1 {
		pxUpdated -= colMax
	} else if pxUpdated < 0 {
		pxUpdated += colMax
	}

	// y-axis/row teleport logic
	if pyUpdated > rowMax-1 {
		pyUpdated -= rowMax
	} else if pyUpdated < 0 {
		pyUpdated += rowMax
	}

	newRobotInfo := RobotInfo{Position{x: pxUpdated, y: pyUpdated}, robotInfo.Velocity}

	return newRobotInfo
}

func countRobotsWithinArea(robotInfoMap map[int]RobotInfo, rowStart int, rowEnd int, colStart int, colEnd int) int {
	count := 0
	for _, robotPosInfo := range robotInfoMap {
		robotPos := robotPosInfo.Position

		isWithinCols := robotPos.x >= colStart && robotPos.x <= colEnd
		isWithinRows := robotPos.y >= rowStart && robotPos.y <= rowEnd

		// increment count if within area
		if isWithinCols && isWithinRows {
			count++
		}
	}

	return count
}

func part1(lines []string) int {
	const (
		ROW_COUNT      = 103
		COL_COUNT      = 101
		NUM_OF_SECONDS = 100
	)

	re := regexp.MustCompile(`-?\d+`)
	robotInfoMap := make(map[int]RobotInfo)

	rows := ROW_COUNT
	cols := COL_COUNT

	// Populate robot map from input
	for i, line := range lines {
		res := toIntArray(re.FindAllString(line, -1))

		px := res[0]
		py := res[1]
		vx := res[2]
		vy := res[3]

		robotDetails := RobotInfo{Position{x: px, y: py}, Velocity{x: vx, y: vy}}
		robotInfoMap[i] = robotDetails
	}

	// Simulate robot movement
	for t := 0; t < NUM_OF_SECONDS; t++ {
		for robotId, robotPosInfo := range robotInfoMap {
			// Update position
			newRobotPosInfo := updateRobotLocation(robotPosInfo, rows, cols)

			// Update map
			robotInfoMap[robotId] = newRobotPosInfo
		}
	}

	// Count robots in each quadrant
	numOfRobotsInTopLeftQuadrant := countRobotsWithinArea(robotInfoMap, 0, (rows/2)-1, 0, (cols/2)-1)
	numOfRobotsInTopRightQuadrant := countRobotsWithinArea(robotInfoMap, 0, (rows/2)-1, (cols/2)+1, cols-1)
	numOfRobotsInBottomLeftQuadrant := countRobotsWithinArea(robotInfoMap, (rows/2)+1, rows-1, 0, (cols/2)-1)
	numOfRobotsInBottomRightQuadrant := countRobotsWithinArea(robotInfoMap, (rows/2)+1, rows-1, (cols/2)+1, cols-1)

	safetyFactor := numOfRobotsInTopLeftQuadrant * numOfRobotsInTopRightQuadrant * numOfRobotsInBottomLeftQuadrant * numOfRobotsInBottomRightQuadrant

	return safetyFactor
}
