package main

import (
	"strconv"
)

/* Part 1 is about iterating through an integer slice for
every "blink", and while iterating, evaluate a set of rules,
applying the first one that matches. Then returning the
number of elements (stones) in the int slice
Rules are as follows:
1. if value 0 replaced with value 1
2. if value has even number of digits, split into two elements,
left taking the left half and right the right half
3. if the above dont apply, multiply value by 2024 */

func solveP2(slice []int) (ret int) {
	const blinkCount = 75
	countMap := make(map[int]int)
	for _, val := range slice {
		countMap[val]++
	}
	for i := 0; i < blinkCount; i++ {
		countMap = blinkFast(countMap)
	}
	total := 0
	for _, count := range countMap {
		total += count
	}
	printIntToIntMap(countMap)
	return total
}

// keep track of the count of the occurrances of every element
func blinkFast(countMap map[int]int) map[int]int {
	newCountMap := make(map[int]int)
	for val, count := range countMap {
		if val == 0 {
			// rule 1
			newCountMap[1] += count
		} else if digitCountFast(val)%2 == 0 {
			// rule 2
			valStr := strconv.Itoa(val)
			mid := len(valStr) / 2
			left, _ := strconv.Atoi(valStr[:mid])
			right, _ := strconv.Atoi(valStr[mid:])
			newCountMap[left] += count
			newCountMap[right] += count
		} else {
			// rule 3
			newCountMap[val*2024] += count
		}
	}
	return newCountMap
}

func digitCountFast(n int) int {
	if n == 0 {
		return 1
	}
	count := 0
	for n != 0 {
		n /= 10
		count++
	}
	return count
}
