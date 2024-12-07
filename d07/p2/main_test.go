package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	// calculate(nums []int, ops []rune) (result int)
	var testCases = []struct {
		name           string
		nums           []int
		ops            []rune
		expectedResult int
	}{
		{"simple add", []int{10, 19}, []rune{'+'}, 29},
		{"simple multiplication", []int{10, 19}, []rune{'*'}, 190},
		{"multiply and add", []int{81, 40, 27}, []rune{'+', '*'}, 3267},
		{"order of operations", []int{11, 6, 16, 20}, []rune{'+', '*', '+'}, 292},
		{"simple concatenation", []int{15, 6}, []rune{'|'}, 156},
		{"multiplication with concatenation", []int{6, 8, 6, 15}, []rune{'*', '|', '*'}, 7290},
		{"addition with concatenation", []int{17, 8, 14}, []rune{'|', '+'}, 192},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := calculate(tc.nums, tc.ops)
			if result != tc.expectedResult {
				t.Errorf("calculate(%v, %v) = %v; want %v", tc.nums, tc.ops, result, tc.expectedResult)
			}
		})
	}
}