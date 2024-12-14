package main

import (
	"fmt"
)

// this helps to find the christmas tree
// completely manually...
// you can skip x moves by typing the number+Enter
func solveP2(slice []string) {
	var height int = 103
	var width int = 101
	var robots []robot = parseRobots(slice)
	var currentMoves = 0
	for {
		robots = moveRobots(robots, height, width)
		currentMoves++
		printRobots(robots, height, width)
		fmt.Printf("%v moves: Does the output look like a Christmas tree? (press y+Enter to terminate, just Enter to continue)\n", currentMoves)
		var input string
		var skip int
		fmt.Scanln(&input)
		if input == "y" {
			break
		}
		if _, err := fmt.Sscanf(input, "%d", &skip); err == nil {
			for i := 1; i < skip; i++ {
				robots = moveRobots(robots, height, width)
				currentMoves++
			}
			continue
		}
	}
	fmt.Printf("Found the Christmas tree in %v moves/seconds\n", currentMoves)
}

func printRobots(robots []robot, height, width int) {
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}

	for _, r := range robots {
		if r.p.y >= 0 && r.p.y < height && r.p.x >= 0 && r.p.x < width {
			grid[r.p.y][r.p.x] = '*'
		}
	}

	for _, row := range grid {
		for _, cell := range row {
			fmt.Print(string(cell))
		}
		fmt.Println()
	}
}
