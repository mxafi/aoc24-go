package main

import (
	"bufio"
	"fmt"
	"os"
)

func readIntGrid(filePath string) ([][]int, error) {
	var file, err = os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("opening input file '%v': %v", filePath, err)
	}
	defer file.Close()

	var grid [][]int
	var scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		var line = scanner.Text()
		var row []int
		for _, char := range line {
			var num = int(char - '0')
			row = append(row, num)
		}
		grid = append(grid, row)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading input file '%v': %v", filePath, err)
	}
	return grid, nil
}

func debugPrintf(format string, v ...interface{}) {
	if globalDebugEnabled {
		fmt.Printf(format, v...)
	}
}

func debugPrintln(v ...interface{}) {
	if globalDebugEnabled {
		fmt.Println(v...)
	}
}

func printIntGrid(grid [][]int) {
	var rowLen int = len(grid[0])
	for i, row := range grid {
		debugPrintf("%v|%v\n", row, i)
	}
	var dashes string
	for i := 0; i < rowLen; i++ {
		dashes += "--"
	}
	dashes += "-"
	debugPrintln(dashes)
	var indexes string
	for i := 0; i < rowLen; i++ {
		indexes += fmt.Sprintf(" %d", i)
	}
	debugPrintln(indexes)
}
