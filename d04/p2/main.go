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

	// simplest search for XMAS or SAMX is character by character, top down
	// where patterns are checked starting with X or S
	// diagonal checks are only done up-right and down-right
	// vertical checks are done only down, and horizontal only left
	for j, row := range grid {
		for i, c := range row {
			if c == 'A' {
				if isCross(grid, c, i, j, rowLen, gridLen) {
					xmasCount += 1
				}
			}
		}
	}

	return xmasCount
}

func isCross(grid [][]rune, c rune, i int, j int, rowLen int, gridLen int) bool {
	// check if cross fits in grid
	if i < 1 || i+2 > rowLen || j < 1 || j+2 > gridLen {
		return false // does not fit, we are at the edge of the grid
	}

	// get the diagonal string and match either one, A in the middle confirmed already
	var topDownRunes = []rune{grid[j-1][i-1], grid[j][i], grid[j+1][i+1]}
	var bottomUpRunes = []rune{grid[j+1][i-1], grid[j][i], grid[j-1][i+1]}
	var topDown string = string(topDownRunes)
	var bottomUp string = string(bottomUpRunes)
	
	if topDown != "MAS" && topDown != "SAM" {
		return false // TD does not match
	}

	if bottomUp != "MAS" && bottomUp != "SAM" {
		return false // BU does not match
	}

	return true
}
