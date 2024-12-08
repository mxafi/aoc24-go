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
			if !isAntennaRune(r) && !isBackgroundRune(r) {
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
	fmt.Println("012345678901")

	fmt.Println("Solving GRID...")

	start := time.Now()

	var uniqueAntinodeLocations = gridSolver(grid)

	elapsed := time.Since(start)
	fmt.Println("Solved in:", elapsed)

	fmt.Println("Unique antinode location count:", uniqueAntinodeLocations)
}

type point struct {
	x int
	y int
}

type frequency struct {
	antennaSymbol  rune    // describes what the symbol to represent an antenna on the frequency is
	antennaPoints  []point // all antenna coordinates for the frequency
	antinodePoints []point // all antinode coordinates for the frequency
}

type limits struct {
	x_min int // minimum x value inclusive
	x_max int // maximum x value exclusive
	y_min int // minimum y value inclusive
	y_max int // maximum y value exclusive
}

func gridSolver(grid [][]rune) (uniqueAntinodeLocations int) {

	// get unique frequencies (a-z, A-Z, 0-9)
	var frequencies = make(map[rune]frequency)
	var uniqFreqCheck = make(map[rune]bool)
	for j, row := range grid {
		for i, r := range row {
			if !isBackgroundRune(r) {
				if !uniqFreqCheck[r] {
					uniqFreqCheck[r] = true
					// initialize the new frequency
					var newFreq = frequency{
						antennaSymbol: r,
						antennaPoints: make([]point, 0),
					}
					newFreq.antennaPoints = append(newFreq.antennaPoints, point{x: i, y: j})
					frequencies[r] = newFreq
				} else {
					// add the antenna point to an existing frequency
					f := frequencies[r]
					f.antennaPoints = append(f.antennaPoints, point{x: i, y: j})
					frequencies[r] = f
				}
			}
		}
	}

	// defining some limits to check against at the time of creating antinodes
	var limit = limits{
		x_min: 0,
		x_max: len(grid[0]),
		y_min: 0,
		y_max: len(grid),
	}

	// calculate the antinodes
	for r, f := range frequencies {
		fmt.Println("calculating antinodes for", string(r))
		var antinodes []point = calculateAntinodes(f.antennaPoints, limit)
		f.antinodePoints = antinodes
		frequencies[r] = f
	}
	fmt.Println("finished all antinode calculations")

	var uniqueAntinodePoints []point
	var uniqAntinodeCheck = make(map[point]bool)
	for _, f := range frequencies {
		for _, an := range f.antinodePoints {
			if !uniqAntinodeCheck[an] {
				uniqAntinodeCheck[an] = true
				if isPointInLimits(an, limit) {
					uniqueAntinodePoints = append(uniqueAntinodePoints, an)
				}
			}
		}
	}
	printPointGrid(uniqueAntinodePoints, limit.y_max, limit.x_max)
	uniqueAntinodeLocations = len(uniqueAntinodePoints)
	return
}

func calculateAntinodes(aps []point, lim limits) (antinodes []point) {
	// antinode exists on every point that is in line with any two
	// antenna points where the other antenna point is twice as far away
	// antinodes are therefore delta vector distance away (x and y)
	// from either antenna point, away from each of the antenna points
	// there are two antinodes for each pair of antenna points

	// for part 2, now there is no distance limits, so antinodes are created
	// every dx dy away from the antenna points

	if len(aps) < 2 {
		fmt.Println("cannot calculate antinodes for less than 2 aps")
	}
	for i, ap1 := range aps {
		for j := i + 1; j < len(aps); j++ {
			// above loops prevent looping through the same pair twice or pairing one ap with itself
			ap2 := aps[j]
			fmt.Printf(" - pair: ap1: (%v,%v), ap2: (%v,%v)\n", ap1.x, ap1.y, ap2.x, ap2.y)
			var dx int = ap2.x - ap1.x
			var dy int = ap2.y - ap1.y
			fmt.Printf("         got dx %v, dy %v\n", dx, dy)

			// going out from ap2, should use +dy +dx.. // edit: reverse for part 2
			for i := 1; ; i++ {
				var antinode = point{x: ap1.x + (dx * i), y: ap1.y + (dy * i)}
				antinodes = append(antinodes, antinode)
				if !isPointInLimits(antinode, lim) {
					break // no more points to add in this direction
				}
			}

			// going out from ap1, should use -dy -dx.. // edit: reverse for part 2
			for i := 1; ; i++ {
				var antinode = point{x: ap2.x - (dx * i), y: ap2.y - (dy * i)}
				antinodes = append(antinodes, antinode)
				if !isPointInLimits(antinode, lim) {
					break // no more points to add in this direction
				}
			}
			fmt.Printf("         got antinodes")
			for _, an := range antinodes {
				fmt.Printf(" (%v,%v)", an.x, an.y)
			}
			fmt.Println()
		}
	}

	return
}

func isAntennaRune(r rune) bool {
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

func isBackgroundRune(r rune) bool {
	return r == '.'
}

func printFrequency(f frequency) {
	fmt.Printf("%v: %v antenna points, %v antinode points\n", string(f.antennaSymbol), len(f.antennaPoints), len(f.antinodePoints))
	fmt.Printf("- ap(s): ")
	for _, point := range f.antennaPoints {
		fmt.Printf("(%v,%v) ", point.x, point.y)
	}
	fmt.Println()
	fmt.Printf("- an(s): ")
	for _, point := range f.antinodePoints {
		fmt.Printf("(%v,%v) ", point.x, point.y)
	}
	fmt.Println()
}

func isPointInLimits(p point, lim limits) bool {
	return p.x >= lim.x_min && p.y >= lim.y_min && p.x < lim.x_max && p.y < lim.y_max
}

// prints a dotted grid with the points represented as hashtags,
// points outside of the grid are hidden
func printPointGrid(points []point, grid_height int, grid_width int) {
	var grid = make([][]rune, grid_height)
	for i := 0; i < grid_height; i++ {
		grid[i] = make([]rune, grid_width)
		for j := 0; j < grid_width; j++ {
			grid[i][j] = '.'
		}
	}
	for _, p := range points {
		grid[p.y][p.x] = '#'
	}
	for i, row := range grid {
		fmt.Printf("%v|%v\n", string(row), i)
	}
	fmt.Println("012345678901")
}
