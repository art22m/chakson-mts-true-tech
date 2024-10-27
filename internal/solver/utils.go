package solver

import "jackson/internal/maze"

func abs(v int) int {
	if v > 0 {
		return v
	}
	return -v
}

func getNearest(x, from, to int) int {
	switch {
	case x < from:
		return from
	case to < x:
		return to
	default:
		return x
	}
}

func getNeighboursWithDirection(pos Position) (res []PositionWithDirection) {
	x, y := pos.x, pos.y
	if validPosition(x-1, y) {
		res = append(res, PositionWithDirection{Position{x - 1, y}, maze.Down})
	}
	if validPosition(x+1, y) {
		res = append(res, PositionWithDirection{Position{x + 1, y}, maze.Up})
	}
	if validPosition(x, y-1) {
		res = append(res, PositionWithDirection{Position{x, y - 1}, maze.Left})
	}
	if validPosition(x, y+1) {
		res = append(res, PositionWithDirection{Position{x, y + 1}, maze.Right})
	}
	return res
}

func getNeighboursNotFinish(pos Position) (res []Position) {
	x, y := pos.x, pos.y
	if checkPositionNotFinish(x-1, y) {
		res = append(res, Position{x - 1, y})
	}
	if checkPositionNotFinish(x+1, y) {
		res = append(res, Position{x + 1, y})
	}
	if checkPositionNotFinish(x, y-1) {
		res = append(res, Position{x, y - 1})
	}
	if checkPositionNotFinish(x, y+1) {
		res = append(res, Position{x, y + 1})
	}
	return res
}

func validPosition(x, y int) bool {
	return 0 <= x && x < height && 0 <= y && y < width
}

func checkPositionNotFinish(x, y int) bool {
	if finishXFrom <= x && x <= finishXTo && finishYFrom <= y && y <= finishYTo {
		return false
	}
	return validPosition(x, y)
}
