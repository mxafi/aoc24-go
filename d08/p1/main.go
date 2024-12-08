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
		for _, r := range row {
			if !isAllowedRune(r) {
				fmt.Printf("grid contains illegal character '%v' in row '%v'\n", string(r), string(row))
				os.Exit(1)
			}
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

	var uniqueAntinodeLocations = gridSolver(grid)

	elapsed := time.Since(start)
	fmt.Println("Solved in:", elapsed)

	fmt.Println("Unique antinode location count:", uniqueAntinodeLocations)
}

func gridSolver(grid [][]rune) (uniqueAntinodeLocations int) {

	// get unique frequencies (a-z, A-Z, 0-9)
	// for every frequency, generate a list of antenna coordinates
	// for every frequency, generate a list of antinode coordinates
	// merge the lists to contain only unique antinode coordinates for all frequencies
	// return the length of the unique merged list

	fmt.Println(len(grid))

	return
}

func isAllowedRune(r rune) bool {
	if r == '.' {
		return true
	}
	if r >= '0' && r <= '9' {
		return true
	}
	if r >= 'a' && r <= 'z' {
		return true
	}
	if r >= 'A' && r <= 'Z' {
		return true
	}
	return false
}
