package main

import (
	"fmt"
	"os"
)

const globalDebugEnabled = true

func main() {
	if len(os.Args) < 3 || len(os.Args) > 3 {
		fmt.Println("usage: ./aoc FILE PART")
		os.Exit(1)
	}
	grid, err := readIntGrid(os.Args[1])
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	printIntGrid(grid)

	var part string = os.Args[2]
	if part == "1" {
		fmt.Println("Part 1: Sum of all trailhead scores:", solveP1(grid))
	} else if part == "2" {
		fmt.Println("Part 2: NOT IMPLEMENTED YET")
	} else {
		fmt.Println("usage: ./aoc FILE PART (PART is 1 or 2)")
		os.Exit(1)
	}
}
