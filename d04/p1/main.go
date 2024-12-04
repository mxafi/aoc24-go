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
		fmt.Printf("%v|%v\n", string(row), i)
	}
	fmt.Println("0123456789")

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
	for j, row := range grid {
		for i, c := range row {
			if c == 'X' || c == 'S' {
				if isVerticalSearchEnabled {
					xmasCount += searchVertical(grid, c, i, j, rowLen, gridLen)
				}
				if isHorizontalSearchEnabled {
					xmasCount += searchHorizontal(grid, c, i, j, rowLen, gridLen)
				}
				if isDiagonalSearchEnabled {
					xmasCount += searchDiagonal(grid, c, i, j, gridLen)
				}
			}
		}
	}

	return xmasCount
}

// searches only top down
// returns count of XMAS or SAMX instances
func searchVertical(grid [][]rune, c rune, i int, j int, rowLen int, gridLen int) int {
	var count int = 0
	if j+4 > gridLen {
		return 0 // cannot fit 4 letter word in grid
	}

	var targetRunes = []rune{grid[j][i], grid[j+1][i], grid[j+2][i], grid[j+3][i]}
	var targetString string = string(targetRunes)

	if targetString == "XMAS" || targetString == "SAMX" {
		count += 1
		fmt.Printf("found vert at (j,i) (%v,%v): %v\n", j, i, targetString)
	}

	return count
}

// searches only left to right
// returns count of XMAS or SAMX instances
func searchHorizontal(grid [][]rune, c rune, i int, j int, rowLen int, gridLen int) int {
	var count int = 0
	if i+4 > gridLen {
		return 0 // cannot fit 4 letter word in grid
	}

	var targetRunes = []rune{grid[j][i], grid[j][i+1], grid[j][i+2], grid[j][i+3]}
	var targetString string = string(targetRunes)

	if targetString == "XMAS" || targetString == "SAMX" {
		count += 1
		fmt.Printf("found hori at (j,i) (%v,%v): %v\n", j, i, targetString)
	}

	return count
}

// searches only left to right (from position to up-right and down-right)
// returns count of XMAS or SAMX instances
func searchDiagonal(grid [][]rune, c rune, i int, j int, gridLen int) int {
	var count int = 0
	if i+4 > gridLen {
		return 0 // cannot fit 4 letter word in grid to right
	}

	// look for up-right first

	if j >= 3 {
		// fits a 4 letter word up
		var targetRunes = []rune{grid[j][i], grid[j-1][i+1], grid[j-2][i+2], grid[j-3][i+3]}
		var targetString string = string(targetRunes)
	
		if targetString == "XMAS" || targetString == "SAMX" {
			count += 1
			fmt.Printf("found diag at (j,i) (%v,%v): %v (  up-right)\n", j, i, targetString)
		}
	}

	// look for down-right next

	if j+4 <= gridLen {
		// fits a 4 letter word down
		var targetRunes = []rune{grid[j][i], grid[j+1][i+1], grid[j+2][i+2], grid[j+3][i+3]}
		var targetString = string(targetRunes)
	
		if targetString == "XMAS" || targetString == "SAMX" {
			count += 1
			fmt.Printf("found diag at (j,i) (%v,%v): %v (down-right)\n", j, i, targetString)
		}
	}

	return count
}
