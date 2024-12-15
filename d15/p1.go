package main

import "log"

func solveP1(input []string) int {
	// parse input
	grid, moves := parseGridAndMoves(input)
	debugPrintln("Input grid:")
	printRuneGrid(grid)
	debugPrintln("Input moves:", string(moves))

	// find robot
	rPos := point{x: -1, y: -1}
	for i, row := range grid {
		for j, c := range row {
			if c == '@' {
				rPos = point{x: j, y: i}
				debugPrintf("Input robot position (%+v)\n", rPos)
			}
		}
	}

	// process moves
	for i, move := range moves {
		rPos = processMove(grid, move, rPos)
		debugPrintln("After move", i+1, ":", string(move))
		printRuneGrid(grid)
	}

	// calculate gps coord sum
	var sum int
	for i, row := range grid {
		for j, c := range row {
			if c == 'O' {
				dTop := i
				dLeft := j
				gpsCoord := 100*dTop + dLeft
				sum += gpsCoord
			}
		}
	}
	return sum
}

type point struct {
	x, y int
}

// rY and rX represent the robot position before processing the move
// returns the new robot position after processing the move
func processMove(grid [][]rune, moveRune rune, rPos point) point {
	var delta point
	switch moveRune {
	case '^':
		// up (towards negative y)
		delta = point{x: 0, y: -1}
	case '>':
		// right
		delta = point{x: 1, y: 0}
	case 'v':
		// down (towards positive y)
		delta = point{x: 0, y: 1}
	case '<':
		// left
		delta = point{x: -1, y: 0}
	default:
		log.Fatal("error: encountered invalid move:", moveRune)
	}

	var rNewPos = rPos
	var firstSpacePos = point{x: rPos.x + delta.x, y: rPos.y + delta.y}
	var firstSpace rune = getGridRune(grid, firstSpacePos)

	switch firstSpace {
	case '#':
		// case 1: space in front of robot is a wall '#', dont move
		return rPos
	case '.':
		// case 2: space in front of robot is free, just move there
		rNewPos = point{x: rPos.x + delta.x, y: rPos.y + delta.y}
		setGridRune(grid, '.', rPos)
		setGridRune(grid, '@', rNewPos)
		return rNewPos
	case 'O':
		// case 3: space in front of robot is a box 'O':
		// continue to case 3.1
	default:
		log.Fatal("error: (case 3) expected 'O', got", string(firstSpace))
	}

	// case 3.1: there is no free space after the box before hitting a wall, dont move
	boxCount := 1
	cPos := point{x: rPos.x + delta.x, y: rPos.y + delta.y}
case31:
	for {
		cPos = point{x: cPos.x + delta.x, y: cPos.y + delta.y}
		switch getGridRune(grid, cPos) {
		case 'O':
			boxCount++
		case '.':
			break case31 // continue to shifting boxes
		case '#':
			return rPos // cannot move anything
		default:
			log.Fatal("error: (case 3.1) unexpected rune:", string(getGridRune(grid, cPos)))
		}
	}

	// shift the boxes into the free space
	setGridRune(grid, '.', rPos) // rPos becomes empty on move
	setGridRune(grid, '@', firstSpacePos) // firstSpacePos becomes robot
	setGridRune(grid, 'O', cPos) // cPos becomes box
	rNewPos = firstSpacePos
	return rNewPos
}

func getGridRune(grid [][]rune, p point) rune {
	return grid[p.y][p.x]
}

func setGridRune(grid [][]rune, r rune, p point) {
	grid[p.y][p.x] = r
}

func parseGridAndMoves(input []string) (grid [][]rune, moves []rune) {
	passedEmptyLine := false
	for _, line := range input {
		if !passedEmptyLine {
			if line == "" {
				passedEmptyLine = true
				continue
			}
			grid = append(grid, []rune(line))
		} else {
			moves = append(moves, []rune(line)...)
		}
	}
	return
}
