package main

/* Part 1 is about detecting all regions in a map,
counting the area and circumference for each region,
multiplying the two, and summing that total. */

func solveP1(grid [][]rune) (ret int) {
	var regions []region

	debugPrintln("Starting to solve grid:")
	printRuneGrid(grid)

	currentRegionId := 0
	visited := make(map[[2]int]bool) // y/i , x/j, isVisited
	for i, row := range grid {
		for j := range row {
			if visited[[2]int{i, j}] {
				continue
			}
			debugPrintf("Exploring region with rune: %s, id: %d, coordinates: (%d, %d)\n", string(grid[i][j]), currentRegionId, i, j)
			reg := exploreRegion(visited, i, j, grid, currentRegionId)
			regions = append(regions, reg)
			currentRegionId++
		}
	}
	debugPrintln("Found region count:", len(regions))

	for _, reg := range regions {
		ret += reg.perimeter * reg.area
	}

	return
}

type region struct {
	char      rune // the character representing the plots in the region
	perimeter int  // contribution to region total perimeter / circumference
	area      int  // count of plots in the region
	id        int  // region id
}

func exploreRegion(visited map[[2]int]bool, i int, j int, grid [][]rune, currentId int) region {
	// we assume i, j is a new region
	char := grid[i][j]
	reg := region{char: char, id: currentId}
	var explore func(int, int) bool // true if explored a plot belonging to the region, false if outside
	explore = func(y, x int) bool {
		if x < 0 || y < 0 || y >= len(grid) || x >= len(grid[0]) {
			return false
		}
		if grid[y][x] != char {
			return false
		}
		if visited[[2]int{y, x}] {
			return true // this is the same char, so it belongs.. but dont recount
		}
		visited[[2]int{y, x}] = true
		reg.area++
		if !explore(y, x-1) {
			reg.perimeter++
		}
		if !explore(y, x+1) {
			reg.perimeter++
		}
		if !explore(y-1, x) {
			reg.perimeter++
		}
		if !explore(y+1, x) {
			reg.perimeter++
		}
		return true
	}
	explore(i, j)
	return reg
}
