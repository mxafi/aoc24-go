package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
)

func getNumbers(line string) (int, int) {
	var leftHalf []rune
	var rightHalf []rune
	var leftFound bool
	var leftDone bool
	var onDigits bool

	for _, r := range line {
		if r >= '0' && r <= '9' {
			if !leftFound {
				leftFound = true
			}
			onDigits = true
			if !leftDone {
				leftHalf = append(leftHalf, r)
			} else {
				rightHalf = append(rightHalf, r)
			}
		} else {
			onDigits = false
		}
		if leftFound && !onDigits {
			leftDone = true
		}
	}
	var leftS = string(leftHalf)
	var rightS = string(rightHalf)

	left, err := strconv.Atoi(leftS)
	if err != nil {
		fmt.Println("error converting string to int:", err)
		os.Exit(1)
	}

	right, err := strconv.Atoi(rightS)
	if err != nil {
		fmt.Println("error converting string to int:", err)
		os.Exit(1)
	}

	return left, right
}

func mergeLists(left []int, right []int) []int {
	var leftLen = len(left)
	var rightLen = len(right)
	if leftLen != rightLen {
		fmt.Println("error: cannot zip slices of different lengths")
		os.Exit(1)
	}

	var merged []int = make([]int, leftLen) // size the new slice correctly ahead of time
	for i := 0; i < leftLen; i++ {
		merged[i] = int(math.Abs(float64(left[i] - right[i])))
	}
	return merged
}

func main() {
	args := os.Args
	var input string = args[1]

	var file, err = os.Open(input)
	if err != nil {
		fmt.Printf("error opening input '%v': %v\n", input, err)
		os.Exit(1)
	}
	defer file.Close()

	var left []int
	var right []int

	var scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		var line = scanner.Text()
		var leftIn, rightIn int = getNumbers(line)
		fmt.Printf("%v:%v\n", leftIn, rightIn)
		left = append(left, leftIn)
		right = append(right, rightIn)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error reading input file:", err)
	}

	fmt.Println("Left  Unsorted List: ", left)
	fmt.Println("Right Unsorted List: ", right)
	slices.Sort(left)
	slices.Sort(right)
	fmt.Println("Left  Sorted   List: ", left)
	fmt.Println("Right Sorted   List: ", right)

	var merged []int = mergeLists(left, right)
	fmt.Println("        Merged List: ", merged)

	var mergedSum int = 0
	for _, val := range merged {
		mergedSum += val
	}

	fmt.Println("    Merged List Sum: ", mergedSum)
}
