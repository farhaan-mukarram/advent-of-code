package main

type Position struct {
	x int
	y int
}

type Velocity struct {
	x int
	y int
}

type RobotInfo struct {
	Position
	Velocity
}
