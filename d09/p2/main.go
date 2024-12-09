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

	var line string

	var scanner = bufio.NewScanner(file)
	var scanRow int = 1
	for scanner.Scan() {
		line = scanner.Text()
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

	// count file and empty block size
	// spread file and empty blocks with their ids
	var memory []int // -1 is empty, 0, 1, ... is file ID
	var currentFileId int = 0
	for i, r := range line {
		// line starts with file, then empty space, etc.
		if r < '0' && r > '9' {
			fmt.Println("invalid characters detected")
			os.Exit(1)
		}
		isFile := false
		if i%2 == 0 {
			isFile = true
		}
		blockCount := int(r - '0')
		for j := 0; j < blockCount; j++ {
			if isFile {
				memory = append(memory, currentFileId)
			} else {
				memory = append(memory, -1)
			}
		}
		if isFile {
			currentFileId++
		}
	}
	printMemory(memory, "input memory")

	// go through the memory back to front, swapping file blocks with empty blocks from the start
	for i := len(memory) - 1; ; i-- {
		if memory[i] < 0 {
			continue // skip empty
		}

		// find out the length of the file
		var fileBlockSize int
		fileId := memory[i]
		for j := i - 1; j >= 0; j-- {
			if memory[j] != fileId {
				fileBlockSize = i - j
				i = j + 1
				break
			}
		}

		// find empty space that fits the file
		var emptyBlockStartIndex int
		var emptyBlockSize int
		for j := 0; j < i; j++ {
			if memory[j] >= 0 {
				continue // skip files
			}
			emptyBlockSize = 0
			for k := j; k < i && memory[k] == -1; k++ {
				emptyBlockSize++
			}
			if emptyBlockSize >= fileBlockSize {
				emptyBlockStartIndex = j
				break
			} else {
				emptyBlockSize = 0
			}
		}

		// check sort finish condition
		if emptyBlockStartIndex >= i {
			break // we're done
		}

		// check current file guard
		if emptyBlockSize == 0 {
			continue // this file cannot fit in any space to its left, skip to next
		}

		// move the file
		for j := 0; j < fileBlockSize; j++ {
			memory[emptyBlockStartIndex+j] = memory[i+j]
			memory[i+j] = -1
		}
		printMemory(memory, "sorting...")
	}

	// calculate the checksum (total from multiplying each index with the value)
	var filesystemChecksum int = 0
	for i, val := range memory {
		if val >= 0 {
			filesystemChecksum += i * val
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Solved in:", elapsed)
	fmt.Println("Filesystem Checksum:", filesystemChecksum)
}

func printMemory(mem []int, msg string) {
	if true {
		return // toggle printMemory on if debugging
	}
	if len(msg) > 20 {
		fmt.Println("error: printMemory: msg too long (>20)")
		os.Exit(1)
	}
	fmt.Printf("%-20s:", msg)
	for _, val := range mem {
		if val == -1 {
			fmt.Print(" .")
		} else {
			fmt.Printf(" %d", val)
		}
	}
	fmt.Println()
}
