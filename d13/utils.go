package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readIntSlice(filePath string) ([]int, error) {
	var file, err = os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("opening input file '%v': %v", filePath, err)
	}
	defer file.Close()

	var slice []int
	var scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		var line = scanner.Text()
		var fields = strings.Fields(line)
		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				return nil, fmt.Errorf("parsing integer from field '%v': %v", field, err)
			}
			slice = append(slice, num)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading input file '%v': %v", filePath, err)
	}
	return slice, nil
}

func printIntSlice(slice []int) {
	if !globalDebugEnabled {
		return
	}
	for _, num := range slice {
		debugPrintf("%d ", num)
	}
	debugPrintln()
}

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

func readRuneGrid(filePath string) ([][]rune, error) {
	var file, err = os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("opening input file '%v': %v", filePath, err)
	}
	defer file.Close()

	var grid [][]rune
	var scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		var line = scanner.Text()
		var row []rune
		for _, char := range line {
			row = append(row, char)
		}
		grid = append(grid, row)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading input file '%v': %v", filePath, err)
	}
	return grid, nil
}

func readStringSlice(filePath string) ([]string, error) {
	var file, err = os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("opening input file '%v': %v", filePath, err)
	}
	defer file.Close()

	var slice []string
	var scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		slice = append(slice, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading input file '%v': %v", filePath, err)
	}
	return slice, nil
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
	if !globalDebugEnabled {
		return
	}
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

func printRuneGrid(grid [][]rune) {
	if !globalDebugEnabled {
		return
	}
	var rowLen int = len(grid[0])
	for i, row := range grid {
		for _, char := range row {
			debugPrintf("%c ", char)
		}
		debugPrintf("|%v\n", i)
	}
	var dashes string
	for i := 0; i < rowLen; i++ {
		dashes += "--"
	}
	dashes += "-"
	debugPrintln(dashes)
	var indexes string
	for i := 0; i < rowLen; i++ {
		indexes += fmt.Sprintf("%d ", i)
	}
	debugPrintln(indexes)
}

func printIntToIntMap(m map[int]int) {
	if !globalDebugEnabled {
		return
	}
	var keys []int
	var values []int
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	for _, key := range keys {
		debugPrintf("%d ", key)
	}
	debugPrintln()
	for i, value := range values {
		debugPrintf("%-*d ", len(fmt.Sprintf("%d", keys[i])), value)
	}
	debugPrintln()
}
