package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	var safeCount int = 0

	var scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		var line = scanner.Text()
		ret, err := isReportSafe(line)
		if err != nil {
			fmt.Println("encountered an error:", err)
			os.Exit(1)
		}
		safeCount += ret
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error reading input file:", err)
	}
	fmt.Println("Number of safe reports:", safeCount)
}

// checks if the report conforms to rules, returns 1 if yes and 0 if not
func isReportSafe(inputReport string) (int, error) {
	report, err := getReportSlice(inputReport)
	if err != nil {
		return 0, err
	}

	// check if already safe
	if isSequenceSafe(report) {
		return 1, nil
	}

	var reportLen int = len(report)

	// check if safe by removing any one level
	for i := 0; i < reportLen; i++ {
		var tempSeq []int = make([]int, reportLen)
		_ = copy(tempSeq, report)
		tempSeq = append(tempSeq[:i], tempSeq[i+1:]...)
		if isSequenceSafe(tempSeq) {
			return 1, nil // report safe with (some) one level removed
		}
	}
	return 0, nil // report unsafe by default
}

func getReportSlice(report string) ([]int, error) {
	var res []int
	var onDigits bool = false
	var digitCollection []rune
	for _, c := range report {
		if c >= '0' && c <= '9' {
			onDigits = true
			digitCollection = append(digitCollection, c)
		} else {
			if onDigits {
				num, err := strconv.Atoi(string(digitCollection))
				if err != nil {
					return nil, fmt.Errorf("atoi failed with: %v", err)
				}
				res = append(res, num)
				digitCollection = digitCollection[:0] // empty without realloc
			}
			onDigits = false
		}
	}
	if onDigits {
		num, err := strconv.Atoi(string(digitCollection))
		if err != nil {
			return nil, fmt.Errorf("atoi failed with: %v", err)
		}
		res = append(res, num)
	}
	return res, nil
}

func isSequenceSafe(seq []int) bool {
	var seqLen int = len(seq)
	if seqLen < 2 {
		return false // seq unsafe
	}

	var directionIncreasing bool = seq[1] > seq[0]
	for i := 1; i < seqLen; i++ {
		var diff int = seq[i] - seq[i-1]

		if (directionIncreasing && diff <= 0) || (!directionIncreasing && diff >= 0) {
			return false // seq unsafe, direction rule violated
		}

		if (directionIncreasing && (diff < 1 || diff > 3)) || (!directionIncreasing && (diff > -1 || diff < -3)) {
			return false // seq unsafe, diff not between 1 and 3 inclusive
		}
	}

	return true // seq safe by default
}
