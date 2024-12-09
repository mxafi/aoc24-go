package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
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

	var scanner = bufio.NewScanner(file)
	var scanRow int = 1
	for scanner.Scan() {
		var line = scanner.Text()
		if scanRow > 1 {
			fmt.Println("more than 1 row exists, invalid input")
			os.Exit(1)
		}
		scanRow++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error reading input file:", err)
	}

	start := time.Now()
	var filesystemChecksum int = 0

	// count file and empty blocks
	// reserve space in rune slice
	// spread file and empty blocks with their ids
	// count empty blocks and record their indexes
	// go through the memory back to front, filling up the empty indexes front to back
	// calculate the checksum

	elapsed := time.Since(start)
	fmt.Println("Solved in:", elapsed)

	fmt.Println("Filesystem Checksum:", filesystemChecksum)
}
