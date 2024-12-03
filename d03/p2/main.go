package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	var input string = args[1]

	var fileContent, err = os.ReadFile(input)
	if err != nil {
		fmt.Printf("error opening input '%v': %v\n", input, err)
		os.Exit(1)
	}

	var content string = string(fileContent)

	var sum int64 = getSum(content)

	fmt.Printf("Sum: %v\n", sum)
}

func getSum(content string) int64 {
	var sum int64 = 0
	var isEnabled bool = true
	var contentLen int = len(content)

	for i, _ := range content {
		if strings.HasPrefix(content[i:], "do()") {
			isEnabled = true
			i += 3
			continue
		}
		if strings.HasPrefix(content[i:], "don't()") {
			isEnabled = false
			i += 6
			continue
		}
		if strings.HasPrefix(content[i:], "mul(") && isEnabled {
			// ensure we have a valid mul, i.e. mul(1,2) with numbers with 1-3 digits
			i += 4 // skip "mul("
			// check elements like (num1, comma, num2, closing bracket) one by one
			// being careful not to run out of bounds
			// the minimum content len after "mul(" should be 4, i.e. "mul(1,2)"
			if i+4 > contentLen {
				return sum
			}

			// setup common variables
			var bracketIdx int
			var commaIdx int

			// depending on if we are at the end, we need two ways to validate the mul
			// max after "mul(" should be 7, i.e. "pxx,yyy)"
			if i+8 >= contentLen {
				// we are at the end and should be careful to not run out of bounds
				bracketIdx = strings.Index(content[i:], ")")
				commaIdx = strings.Index(content[i:], ",")
			} else {
				// we are not near the end and can use fixed length checks to optimise
				bracketIdx = strings.Index(content[i:i+8], ")")
				commaIdx = strings.Index(content[i:i+4], ",")
			}
			if bracketIdx == -1 {
				continue // invalid mul
			}
			if commaIdx == -1 {
				continue // invalid mul
			}
			// validate comma and closing bracket are in proper order
			if commaIdx > bracketIdx {
				continue // invalid mul
			}
			var num1Str string = content[i : i+commaIdx]
			var num2Str string = content[i+commaIdx+1 : i+bracketIdx]
			// validate numstrings, make the nums, multiply and add to sum
			if !isValidNumber(num1Str) || !isValidNumber(num2Str) {
				continue // invalid mul
			}
			num1, _ := strconv.ParseInt(num1Str, 10, 64)
			num2, _ := strconv.ParseInt(num2Str, 10, 64)
			sum += num1 * num2
		}
	}

	return sum
}

func isValidNumber(numStr string) bool {
	// check if numStr is between 1 and 3 digits
	if len(numStr) < 1 || len(numStr) > 3 {
		return false
	}
	// check if numStr contains only digits
	for _, c := range numStr {
		if !isDigit(c) {
			return false
		}
	}
	return true
}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}
