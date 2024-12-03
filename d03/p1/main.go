package main

import (
	"bufio"
	"fmt"
	"os"
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
	for scanner.Scan() {
		var line = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error reading input file:", err)
	}
}
