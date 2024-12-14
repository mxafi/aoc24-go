package main

import (
	"fmt"
	"os"
	"time"
)

const globalDebugEnabled = false

func main() {
	if globalDebugEnabled {
		fmt.Println("Debug messaged are enabled")
	} else {
		fmt.Println("Debug messaged are disabled (better for true input)")
	}
	if len(os.Args) < 3 || len(os.Args) > 3 {
		fmt.Println("usage: ./aoc FILE PART")
		os.Exit(1)
	}
	slice, err := readStringSlice(os.Args[1])
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	
	start := time.Now()
	var part string = os.Args[2]
	if part == "1" {
		fmt.Println("Part 1: Safety factor after 100 seconds:", solveP1(slice))
	} else if part == "2" {
		// todo
	} else {
		fmt.Println("usage: ./aoc FILE PART (PART is 1 or 2)")
		os.Exit(1)
	}
	elapsed := time.Since(start)
	fmt.Println("Elapsed time:", elapsed)
}
