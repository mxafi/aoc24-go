package main

import (
	"bufio"
	"fmt"
	"os"
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

	var xmasCount int = 0
	var grid [][]rune
	var row []rune
	var rowLen int = 0

	var scanner = bufio.NewScanner(file)
	var scanRow int = 1
	for scanner.Scan() {
		var line = scanner.Text()
		row = []rune(line)
		if rowLen == 0 {
			rowLen = len(row)
			if rowLen < 1 {
				fmt.Println("first row empty")
			}
		}
		if rowLen != len(row) {
			fmt.Printf("row %v inconsistent with previous rows (%v): %v\n", scanRow, rowLen, len(row))
		}
		grid = append(grid, row)
		scanRow++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error reading input file:", err)
	}
	fmt.Println("Got GRID")
	fmt.Println("len:", rowLen)
	for i, row := range grid {
		fmt.Printf("%v|%v\n", row, i+1)
	}

	fmt.Println("Solving GRID...")
	xmasCount = solveGrid(grid)

	fmt.Printf("Sum: %v\n", xmasCount)
}

// returns the number of XMAS instances
func solveGrid(grid [][]rune) int {
	var xmasCount int = 0
	var rowLen = len(grid[0])
	var gridLen = len(grid)
	var isVerticalSearchEnabled bool = true
	var isHorizontalSearchEnabled bool = true
	var isDiagonalSearchEnabled bool = true

	// disable search types based on grid shape
	// XMAS fits horizontally and diagonally in 4 spaces
	if rowLen < 4 {
		isHorizontalSearchEnabled = false
		isDiagonalSearchEnabled = false
	}
	// XMAS fits vertically and diagonally in 4 spaces
	if gridLen < 4 {
		isVerticalSearchEnabled = false
		isDiagonalSearchEnabled = false
	}
	// now only the possible search types are enabled

	// simplest search for XMAS or SAMX is character by character, top down
	// where patterns are checked starting with X or S
	// diagonal checks are only done up-right and down-right
	// vertical checks are done only down, and horizontal only left
	for i, row := range grid {
		for j := range row {
			if isVerticalSearchEnabled {
				xmasCount += searchVertical(grid, i, j)
			}
			if isHorizontalSearchEnabled {
				xmasCount += searchHorizontal(grid, i, j)
			}
			if isDiagonalSearchEnabled {
				xmasCount += searchDiagonal(grid, i, j)
			}
		}
	}

	return xmasCount
}

// searches only top down
// returns count of XMAS or SAMX instances
func searchVertical(grid [][]rune, i int, j int) int {
	var count int = 0
	return count
}

// searches only left to right
// returns count of XMAS or SAMX instances
func searchHorizontal(grid [][]rune, i int, j int) int {
	var count int = 0
	return count
}

// searches only left to right (from position to up-right and down-right)
// returns count of XMAS or SAMX instances
func searchDiagonal(grid [][]rune, i int, j int) int {
	var count int = 0
	return count
}
