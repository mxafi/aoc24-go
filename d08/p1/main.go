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

	// calculate the antinodes
	for r, f := range frequencies {
		fmt.Println("calculating antinodes for", string(r))
		var antinodes []point = calculateAntinodes(f.antennaPoints)
		f.antinodePoints = antinodes
		frequencies[r] = f
	}
	fmt.Println("finished all antinode calculations")

	var uniqueAntinodes = make(map[point]bool)
	for _, f := range frequencies {
		for _, an := range f.antinodePoints {
			if !uniqueAntinodes[an] {
				uniqueAntinodes[an] = true
				if an.x >= 0 && an.y >= 0 && an.x < len(grid[0]) && an.y < len(grid) {
					uniqueAntinodeLocations++
				}
			}
		}
	}

	return
}

func calculateAntinodes(aps []point) (antinodes []point) {
	// antinode exists on every point that is in line with any two
	// antenna points where the other antenna point is twice as far away
	// antinodes are therefore delta vector distance away (x and y)
	// from either antenna point, away from each of the antenna points
	// there are two antinodes for each pair of antenna points

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

			// going out from ap2, should use +dy +dx
			var an1 = point{x: ap2.x + dx, y: ap2.y + dy}

			// going out from ap1, should use -dy -dx
			var an2 = point{x: ap1.x - dx, y: ap1.y - dy}

			fmt.Printf("         got antinodes (%v,%v) and (%v,%v)\n", an1.x, an1.y, an2.x, an2.y)
			antinodes = append(antinodes, an1, an2)
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
