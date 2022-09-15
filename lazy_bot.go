package main

import (
	"math"
)

const (
	MoveNorth = "N"
	MoveSouth = "S"
	MoveEast  = "E"
	MoveWest  = "W"

	DropPizza = "D"
)

func convertDeltaToPath(delta int, direct, opposite string) string {
	var (
		direction, path string
	)

	direction = direct
	if delta < 0 {
		direction = opposite
	}

	deltaAbs := int(math.Abs(float64(delta)))
	for i := 0; i < deltaAbs; i++ {
		path = path + direction
	}

	return path
}

// buildLazyPath returns path for bot.
// path is not optimal, path is very lazy.
// the main idea: go via a lane and deliver pizza in each house on the lane
//
// math supporting: we calculate distance between two points
// and convert these distances (for X and Y separately) to path legend.
// when we deliver pizza update our "start" position to move to the next receiver.
// the general principle of moving: we always move towards E-direction
func buildLazyPath(districtScheme DistrictSchemeType) string {
	var (
		x, y, deltaX, deltaY int
		path                 string
	)

	for _, lane := range districtScheme.laneList {
		for _, house := range districtScheme.receivers[lane] {
			deltaX = lane - x
			x = lane

			deltaY = house - y
			y = house

			path = path +
				convertDeltaToPath(deltaX, MoveEast, MoveWest) +
				convertDeltaToPath(deltaY, MoveNorth, MoveSouth) +
				DropPizza
		}
	}

	return path
}
