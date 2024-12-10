package main

/* Part 1 is about finding all 0 values "trailheads" from the grid
then traversing the grid in horizontal and vertical directions where
the value increases by 1 compared to the current one...
Every "trail" found that leads to a 9 is a valid "trail".
A "trailhead" is scored based on the amount of "trails"
that lead to a unique 9 location for that "trailhead",
giving 1 score per reached 9.
The goal is to find out the score for all "trailheads",
and sum the score of each one. */

func solveP1(grid [][]int) (ret int) {
	var trailheadPoints []point = getTrailheadPoints(grid)
	debugPrintln("Got trailheadPoints:", trailheadPoints)
	for _, p := range trailheadPoints {
		var finishPoints []point = findTrails(grid, p)
		debugPrintf("For trailhead (%v) got finishPoints: %v\n", p, finishPoints)
		var unique = filterUniquePoints(finishPoints)
		debugPrintf("For trailhead (%v) got unique: %v\n", p, unique)
		ret += len(unique)
	}
	return
}

func findTrails(grid [][]int, start point) (finishPoints []point) {

	var helper func(pos point)
	helper = func(pos point) {
		directions := []point{{x: 0, y: 1}, {x: 1, y: 0}, {x: 0, y: -1}, {x: -1, y: 0}}
		for _, dir := range directions {
			newPos := point{x: pos.x + dir.x, y: pos.y + dir.y}
			if newPos.y >= 0 && newPos.y < len(grid) && newPos.x >= 0 && newPos.x < len(grid[0]) {
				if grid[newPos.y][newPos.x] == grid[pos.y][pos.x]+1 {
					if grid[newPos.y][newPos.x] == 9 {
						finishPoints = append(finishPoints, newPos)
					} else {
						helper(newPos)
					}
				}
			}
		}
	}

	helper(start)

	return
}

func filterUniquePoints(input []point) (unique []point) {
	seen := make(map[point]bool)
	for _, p := range input {
		if !seen[p] {
			unique = append(unique, p)
			seen[p] = true
		}
	}
	return
}

type point struct {
	x int
	y int
}

func getTrailheadPoints(grid [][]int) (trailheadPoints []point) {
	for i, row := range grid {
		for j, c := range row {
			if c == 0 {
				trailheadPoints = append(trailheadPoints, point{x: j, y: i})
			}
		}
	}
	return
}
