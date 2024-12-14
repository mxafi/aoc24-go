package main

import (
	"fmt"
	"os"
)

type robot struct {
	p point // position
	v point // velocity
}

type point struct {
	x, y int
}

func solveP1(slice []string) int {
	var height int = 103
	var width int = 101
	var robots []robot = parseRobots(slice)
	const moveCountTarget int = 100
	for i := 0; i < moveCountTarget; i++ {
		debugPrintln("starting move", i+1)
		robots = moveRobots(robots, height, width)
	}
	// count number of robots in the quadrants (excl. middle)
	var q1, q2, q3, q4 int // top left, top right, bottom right, bottom left
	var midWidth int = width / 2
	var midHeight int = height / 2
	for i, r := range robots {
		debugPrintf("counting robot %d at position %+v\n", i, r.p)
		if r.p.x >= 0 && r.p.x < midWidth {
			// left side
			debugPrintln("  left side")
			if r.p.y >= 0 && r.p.y < midHeight {
				// top left quadrant q1
				debugPrintln("  in q1")
				q1++
			} else if r.p.y > midHeight && r.p.y < height {
				// bottom left quadrant q4
				debugPrintln("  in q4")
				q4++
			}
		} else if r.p.x > midWidth && r.p.x < width {
			// right side
			debugPrintln("  right side")
			if r.p.y >= 0 && r.p.y < midHeight {
				// top right quadrant q2
				debugPrintln("  in q2")
				q2++
			} else if r.p.y > midHeight && r.p.y < height {
				// bottom right quadrant q3
				debugPrintln("  in q3")
				q3++
			}
		}
	}
	debugPrintf("q1: %d, q2: %d, q3: %d, q4: %d\n", q1, q2, q3, q4)
	// multiply the counts to get safety factor to return
	return q1 * q2 * q3 * q4
}

func parseRobots(lines []string) []robot {
	var robots []robot
	for i, line := range lines {
		var r robot
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &r.p.x, &r.p.y, &r.v.x, &r.v.y)
		robots = append(robots, r)
		debugPrintf("parsed robot %v: %+v\n", i, r)
	}
	return robots
}

func moveRobots(robots []robot, height int, width int) []robot {
	for i, r := range robots {
		robots[i] = moveRobot(r, height, width)
		debugPrintf("robot %d moved to: %+v\n", i, robots[i])
	}
	return robots
}

func moveRobot(r robot, h int, w int) robot {
	// robots teleport between left and right edges for x
	// top and bottom edges for y
	r.p.x = wrapValue(r.p.x, r.v.x, w)
	r.p.y = wrapValue(r.p.y, r.v.y, h)
	return r
}

// n is the current value, d is the delta change for the value, w is the wrap length
// returns the new value
func wrapValue(n int, d int, w int) int {
	if n < 0 || n >= w {
		debugPrintf("error: wrapValue: value already out of range 0-%v: %v\n", w, n)
		os.Exit(1)
	}
	result := (n + d) % w
	if result < 0 {
		result += w
	}
	return result
}
