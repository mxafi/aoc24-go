package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
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

	var cache = make(map[int]int)

	var merged []int = make([]int, leftLen) // size the new slice correctly ahead of time
	for i := 0; i < leftLen; i++ {
		var leftVal int = left[i]
		if value, exists := cache[leftVal]; exists {
			merged[i] = value
		} else {
			var res int = leftVal * countElementInSortedSlice(right, leftVal, rightLen)
			cache[leftVal] = res
			merged[i] = res
		}
	}
	return merged
}

func countElementInSortedSlice(slice []int, target int, sliceLen int) int {
	var firstIdx int = sort.Search(sliceLen, func(i int) bool {
		return slice[i] >= target
	})
	if firstIdx == sliceLen || slice[firstIdx] != target {
		return 0
	}
	var lastIdx int = sort.Search(sliceLen, func(i int) bool {
		return slice[i] > target
	}) - 1
	return lastIdx - firstIdx + 1
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
		left = append(left, leftIn)
		right = append(right, rightIn)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error reading input file:", err)
	}

	slices.Sort(left)
	slices.Sort(right)
	var merged []int = mergeLists(left, right)
	var mergedSum int = 0
	for _, val := range merged {
		mergedSum += val
	}

	fmt.Println("Similarity Score: ", mergedSum)
}
