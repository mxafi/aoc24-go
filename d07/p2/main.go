package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type equation struct {
	result  int
	numbers []int
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

	var inputEquations []equation
	var sum int = 0

	var scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		var line = scanner.Text()
		var lineEq = parseEquation(line)
		inputEquations = append(inputEquations, lineEq)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error reading input file:", err)
	}

	start := time.Now()
	sum = testEquations(inputEquations)
	elapsed := time.Since(start)

	fmt.Printf("Sum: %v\n", sum)
	fmt.Printf("Execution time: %s\n", elapsed)
}

// test every equation in the given slice, trying to insert + and * operators to get the result,
// for every success, the return value is increased
func testEquations(input []equation) (output int) {
	for _, eq := range input {
		var solutionFound bool = false
		var operatorPermutations [][]rune = generatePermutations(len(eq.numbers) - 1)
		var res int
		for _, perm := range operatorPermutations {
			res = calculate(eq.numbers, perm)
			if res == eq.result {
				solutionFound = true
				break
			}
		}
		if solutionFound {
			output += res
		}
	}
	return
}

// generates all possible permutations with + and * operators for some given length
func generatePermutations(length int) (permutations [][]rune) {
	if length < 1 {
		fmt.Println("gen error: cannot generate for len < 1")
		os.Exit(1)
	}

	var helper func([]rune, int)
	helper = func(current []rune, remaining int) {
		if remaining == 0 {
			// found a permutation for the given length
			perm := make([]rune, len(current))
			copy(perm, current)
			permutations = append(permutations, perm)
			return
		}
		// fork recursion for every possible operator
		helper(append(current, '+'), remaining-1)
		helper(append(current, '*'), remaining-1)
		helper(append(current, '|'), remaining-1)
	}

	helper([]rune{}, length)
	return
}

// calculates the result for some list of numbers and operators
// evaluates from left to right, not caring for precedence rules
func calculate(nums []int, ops []rune) (result int) {
	if len(nums) != (len(ops) + 1) {
		fmt.Println("calc error: wrong amount of numbers or operators")
		os.Exit(1)
	}
	var num1 int
	var num2 int
	for i, op := range ops {
		num1 = result
		if i == 0 {
			num1 = nums[i]
		}
		num2 = nums[i+1]
		switch op {
		case '+':
			result = (num1 + num2)
		case '*':
			result = (num1 * num2)
		case '|':
			// concat e.g. 5 | 2 = 52
			s1 := strconv.Itoa(num1)
			s2 := strconv.Itoa(num2)
			concat := s1 + s2
			var err error
			result, err = strconv.Atoi(concat)
			if err != nil {
				fmt.Println("calc error: concat atoi fail")
				os.Exit(1)
			}
		default:
			fmt.Println("calc error: wrong operator symbol")
			os.Exit(1)
		}
	}
	return
}

func parseEquation(line string) equation {
	parts := strings.Split(line, ": ")
	if len(parts) != 2 {
		fmt.Printf("invalid equation format: %v\n", line)
		os.Exit(1)
	}

	result, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Printf("invalid result value: %v\n", parts[0])
		os.Exit(1)
	}

	numberStrings := strings.Split(parts[1], " ")
	numbers := make([]int, len(numberStrings))
	for i, numStr := range numberStrings {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Printf("invalid number value: %v\n", numStr)
			os.Exit(1)
		}
		numbers[i] = num
	}

	return equation{
		result:  result,
		numbers: numbers,
	}
}
