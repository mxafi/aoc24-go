package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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
	var sum int64 = 0
	var scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		var line = scanner.Text()
		var lineSum int64 = getLineSum(line)
		sum += lineSum
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error reading input file:", err)
	}

	fmt.Printf("Sum: %v\n", sum)
}

func getLineSum(line string) int64 {
	var sum int64 = 0

	expr := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := expr.FindAllString(line, -1)

	for _, match := range matches {
		var num1, num2 int64 = getNumbers(match)
		sum += num1 * num2
	}

	return sum
}

func getNumbers(line string) (int64, int64) {
	var num1 int64
	var num2 int64
	var comma = strings.Index(line, ",")
	var num1Str = line[4:comma]
	var num2Str = line[comma+1 : len(line)-1]
	var err error
	num1, err = strconv.ParseInt(num1Str, 10, 64)
	if err != nil {
		fmt.Println("error converting string to int:", err)
		os.Exit(1)
	}
	num2, err = strconv.ParseInt(num2Str, 10, 64)
	if err != nil {
		fmt.Println("error converting string to int:", err)
		os.Exit(1)
	}
	return num1, num2
}
