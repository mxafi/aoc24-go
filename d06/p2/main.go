package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	args := os.Args
	var input string = args[1]

	var file, err = os.Open(input)
	if err != nil {
		fmt.Printf("error opening input '%v': %v\n", input, err)
		os.Exit(1)
	}
	defer file.Close()

	var grid [][]rune
	var row []rune

	var scanner = bufio.NewScanner(file)
	var scanRow int = 1
	for scanner.Scan() {
		var line = scanner.Text()
		row = []rune(line)
		grid = append(grid, row)
		scanRow++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error reading input file:", err)
	}
	if len(grid) == 0 {
		fmt.Println("empty grid")
		os.Exit(1)
	}
	fmt.Println("Got GRID")
	var rowLen int = len(grid[0])
	for i, row := range grid {
		if rowLen != len(row) {
			fmt.Println("inconsistent row lengths")
			os.Exit(1)
		}
		fmt.Printf("%v|%v\n", string(row), i)
	}
	if rowLen == 0 {
		fmt.Println("empty rows")
		os.Exit(1)
	}
	fmt.Println("0123456789")

	fmt.Println("Solving GRID...")

	start := time.Now()
	loopCount, _ := solveGrid(grid, 0)
	elapsed := time.Since(start)
	fmt.Println("Loop count:", loopCount)
	fmt.Printf("Execution time: %s\n", elapsed)
}

type guard struct {
	x  int // current x position
	y  int // current y position
	dx int // current x movement amplitude ( 0 at the beginning)
	dy int // current y movement amplitude (-1 at the beginning, indicating up)
	sx int // starting x position
	sy int // starting x position
}

func solveGrid(grid [][]rune, depth int) (loopCount int, stuckLoop bool) {
	// row and grid length is invariants
	var rowLen int = len(grid[0])
	var gridLen int = len(grid)

	var guard = guard{
		x:  -1, // -1 means not found
		y:  -1, // ...
		dx: 0,  // guard moves up at the beginning
		dy: -1, // ...
		sx: -1,
		sy: -1,
	}

	for j, row := range grid {
		for i, c := range row {
			// search for the guard start position
			if c == '^' {
				guard.x = i
				guard.y = j
				guard.sx = i
				guard.sy = j
			}
		}
	}
	if guard.x == -1 {
		fmt.Println("no guard position found")
		os.Exit(1)
	}

	var maxMoves int = gridLen * rowLen
	var moveCount int = 0
	for {
		var outOfBounds bool = moveGuard(&guard, grid)
		moveCount++
		if outOfBounds {
			break // finished no loop
		}
		if moveCount > maxMoves {
			stuckLoop = true
			loopCount = 1
			return
		}
	}

	if depth == 0 {
		for j, row := range grid {
			for i := range row {
				if j == guard.sy && i == guard.sx {
					continue
				}
				tmpGrid := copyGrid(grid)
				tmpGrid[j][i] = '#'
				if count, stuck := solveGrid(tmpGrid, depth+1); stuck {
					loopCount += count
				}
			}
		}
	}

	return
}

func moveGuard(guard *guard, grid [][]rune) bool {
	var rowLen int = len(grid[0])
	var gridLen int = len(grid)
	var newX int = guard.x + guard.dx
	var newY int = guard.y + guard.dy

	if newX >= rowLen || (guard.x == 0 && guard.dx == -1) || newY >= gridLen || (guard.y == 0 && guard.dy == -1) {
		return true // new position would be out of bounds
	}

	var obsticleCheckCount int = 0
	for {
		if grid[newY][newX] == '#' {
			// found obsticle at new position, should change direction
			rotateGuardCW(guard)
			newX = guard.x + guard.dx
			newY = guard.y + guard.dy
			obsticleCheckCount++
			if newX >= rowLen || (guard.x == 0 && guard.dx == -1) || newY >= gridLen || (guard.y == 0 && guard.dy == -1) {
				return true // new position would be out of bounds
			}
		} else {
			break
		}
		if obsticleCheckCount > 4 {
			fmt.Println("error: guard trapped with obsticles")
			os.Exit(1)
		}
	}

	// move and mark as visited
	guard.x = newX
	guard.y = newY

	return false // is guard out of bounds?
}

func rotateGuardCW(guard *guard) {
	switch {
	case guard.dx == 0 && guard.dy == -1: // up to right
		guard.dx = 1
		guard.dy = 0
	case guard.dx == 1 && guard.dy == 0: // right to down
		guard.dx = 0
		guard.dy = 1
	case guard.dx == 0 && guard.dy == 1: // down to left
		guard.dx = -1
		guard.dy = 0
	case guard.dx == -1 && guard.dy == 0: // left to up
		guard.dx = 0
		guard.dy = -1
	}
}

func copyGrid(grid [][]rune) [][]rune {
	copyGrid := make([][]rune, len(grid))
	for i, row := range grid {
		copyRow := make([]rune, len(row))
		copy(copyRow, row)
		copyGrid[i] = copyRow
	}
	return copyGrid
}

func printGrid[T int | rune](grid [][]T, msg string) {
	fmt.Println("Printing Grid:", msg)
	for i, row := range grid {
		for _, elem := range row {
			if v, ok := any(elem).(rune); ok {
				fmt.Printf("%c", v)
			} else {
				fmt.Printf("%v", elem)
			}
		}
		fmt.Printf("|%v\n", i)
	}
	fmt.Println("0123456789")
}

func printGuard(g guard, msg string) {
	fmt.Printf("guard at (%v,%v) with direction (%v,%v): %v\n", g.x, g.y, g.dx, g.dy, msg)
}
