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

func solveP1(slice []int) (ret int) {
	const blinkCount = 25
	for i := 0; i < blinkCount; i++ {
		slice = blink(slice)
	}
	return len(slice)
}

// will iterate through the elements and apply the rules
func blink(slice []int) []int {
	var newSlice []int
	for i := 0; i < len(slice); i++ {
		var val int = slice[i]
		if val == 0 {
			// rule 1
			newSlice = append(newSlice, 1)
		} else if digitCount(val)%2 == 0 {
			// rule 2
			var valStr = strconv.Itoa(val)
			mid := len(valStr) / 2
			left, _ := strconv.Atoi(valStr[:mid])
			right, _ := strconv.Atoi(valStr[mid:])
			newSlice = append(newSlice, left, right)
		} else {
			// rule 3
			newSlice = append(newSlice, val*2024)
		}
	}
	printIntSlice(newSlice)
	return newSlice
}

func digitCount(n int) int {
	return len(strconv.Itoa(n))
}
