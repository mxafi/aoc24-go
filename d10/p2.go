package main

func solveP2(grid [][]int) (ret int) {
	var trailheadPoints []point = getTrailheadPoints(grid)
	debugPrintln("Got trailheadPoints:", trailheadPoints)
	for _, p := range trailheadPoints {
		var finishPoints []point = findTrails(grid, p)
		debugPrintf("For trailhead (%v) got finishPoints: %v\n", p, finishPoints)
		ret += len(finishPoints)
	}
	return
}