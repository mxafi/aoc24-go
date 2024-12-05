package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// something like 55,11,69
// (has to follow relevant rules and have odd number of pages)
type update struct {
	pages []int
}

// something like 55|69
// (specifies two pages, where the first must appear before the second)
type rule struct {
	pageA int
	pageB int
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

	var sum int64 = 0

	// going to parse rules first, then updates, seperated by single empty line
	var emptyLineDetected bool = false
	var rulesInput []string
	var updatesInput []string

	var scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		var line = scanner.Text()
		if line == "" {
			emptyLineDetected = true
			continue
		}
		if !emptyLineDetected {
			// collect into rulesInput
			rulesInput = append(rulesInput, line)
		} else {
			// collect into updatesInput
			updatesInput = append(updatesInput, line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error reading input file:", err)
	}

	var rules []rule = parseRules(rulesInput)
	var updates []update = parseUpdates(updatesInput)

	// graph pages to point to slice of pages that come after it
	var ruleGraph = make(map[int][]int)
	for _, rule := range rules {
		ruleGraph[rule.pageA] = append(ruleGraph[rule.pageA], rule.pageB)
	}

	for _, currentUpdate := range updates {
		if isUpdateValid(currentUpdate, rules) {
			// fmt.Printf(" valid  update: %v\n", currentUpdate)
		} else {
			fmt.Printf("invalid update: %v\n", currentUpdate)
			var fixed update = fixUpdate(currentUpdate, ruleGraph)
			var midSum int64 = int64(fixed.pages[(len(fixed.pages)-1)/2])
			fmt.Printf("  fixed update: %v (mid:%v)\n", fixed, midSum)
			sum += midSum
		}
	}

	fmt.Printf("Sum: %v\n", sum)
}

func parseRules(input []string) []rule {
	var rules = make([]rule, 0, len(input))
	for _, ruleString := range input {
		if len(ruleString) < 3 {
			fmt.Println("error parsing rule (len)  :", ruleString)
			os.Exit(1)
		}
		split := strings.Split(ruleString, "|")
		if len(split) != 2 {
			fmt.Println("error parsing rule (split):", ruleString)
			os.Exit(1)
		}
		for _, c := range split[0] {
			if c < '0' || c > '9' {
				fmt.Println("error parsing rule (digit-s0):", ruleString)
				os.Exit(1)
			}
		}
		for _, c := range split[1] {
			if c < '0' || c > '9' {
				fmt.Println("error parsing rule (digit-s1):", ruleString)
				os.Exit(1)
			}
		}
		pageA, err := strconv.Atoi(split[0])
		if err != nil {
			fmt.Println("error parsing rule (atoi-s0):", ruleString)
			os.Exit(1)
		}
		pageB, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Println("error parsing rule (atoi-s1):", ruleString)
			os.Exit(1)
		}
		var newRule = rule{pageA: pageA, pageB: pageB}
		rules = append(rules, newRule)
	}
	return rules
}

func parseUpdates(input []string) []update {
	var updates = make([]update, 0, len(input))
	for _, updateString := range input {
		pages := strings.Split(updateString, ",")
		var newUpdate = update{
			pages: make([]int, 0, len(pages)),
		}
		if len(pages) < 3 {
			fmt.Println("error parsing update (split):", updateString)
			os.Exit(1)
		}
		if len(pages)%2 == 0 {
			fmt.Println("error parsing update (not-odd):", updateString)
			os.Exit(1)
		}
		for _, page := range pages {
			for _, c := range page {
				if c < '0' || c > '9' {
					fmt.Println("error parsing update (digits):", updateString)
					os.Exit(1)
				}
			}
			pageInt, err := strconv.Atoi(page)
			if err != nil {
				fmt.Println("error parsing update (atoi):", updateString)
				os.Exit(1)
			}
			newUpdate.pages = append(newUpdate.pages, pageInt)
		}
		updates = append(updates, newUpdate)
	}
	return updates
}

func isUpdateValid(update update, rules []rule) bool {
	pageIndices := make(map[int]int)
	for i, page := range update.pages {
		pageIndices[page] = i
	}

	for _, rule := range rules {
		aIndex, aExists := pageIndices[rule.pageA]
		bIndex, bExists := pageIndices[rule.pageB]
		if aExists && bExists && aIndex > bIndex {
			return false
		}
	}
	return true
}

func fixUpdate(input update, ruleGraph map[int][]int) update {
	// DFS of all edges
	visitedEdges := make(map[int]map[int]bool)
	result := []int{}
	var visit func(int)
	visit = func(page int) {
		if visitedEdges[page] == nil {
			visitedEdges[page] = make(map[int]bool)
		}
		for _, nextPage := range ruleGraph[page] {
			if !visitedEdges[page][nextPage] {
				visitedEdges[page][nextPage] = true
				visit(nextPage)
			}
		}
		result = append(result, page)
	}
	for _, page := range input.pages {
		visit(page)
	}

	// DFS returns results in reverse order, unreverse
	reversed := make([]int, len(result))
	for i, page := range result {
		reversed[len(result)-1-i] = page
	}

	// get the correct count per page to prevent removing duplicates
	pageCount := make(map[int]int)
	for _, page := range input.pages {
		pageCount[page]++
	}

	var fixedPages []int
	for _, page := range reversed {
		if pageCount[page] > 0 {
			fixedPages = append(fixedPages, page)
			pageCount[page]--
		}
	}

	return update{pages: fixedPages}
}
