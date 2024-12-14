package main

import (
	"fmt"
	"testing"
)

func TestWrapValue(t *testing.T) {
	tests := []struct {
		n, d, w, expected int
	}{
		{0, 1, 10, 1},
		{9, 1, 10, 0},
		{0, -1, 10, 9},
		{1, -2, 10, 9},
		{5, 5, 10, 0},
		{5, -5, 10, 0},
		{5, 15, 10, 0},
		{5, -15, 10, 0},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("n=%d,d=%d,w=%d", tt.n, tt.d, tt.w), func(t *testing.T) {
			result := wrapValue(tt.n, tt.d, tt.w)
			if result != tt.expected {
				t.Errorf("wrapValue(%d, %d, %d) = %d; want %d", tt.n, tt.d, tt.w, result, tt.expected)
			}
		})
	}
}
