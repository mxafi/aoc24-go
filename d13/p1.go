package main

import (
	"container/list"
	"regexp"
	"strconv"
)

func solveP1(slice []string) (ret int) {
	machines := parseInput(slice)
	for _, m := range machines {
		mTok := solveMachine(m)
		debugPrintf("  used %v tokens to win\n", mTok)
		ret += mTok
	}

	return
}

// returns the token cost of getting the prize with the least tokens
// token cost is 0 for impossible machines
// token cost is calculated by taking the least moves of
// a times 3, added to the least moves of b
func solveMachine(m machine) (leastTokensToWin int) {
	debugPrintf("solving machine with a(%v,%v), b(%v,%v), p(%v,%v)\n", m.a.x, m.a.y, m.b.x, m.b.y, m.prize.x, m.prize.y)
	start := sPos{x: 0, y: 0, a_steps: 0, b_steps: 0}
	visited := make(map[[2]int]bool)
	visited[[2]int{0, 0}] = true
	lst := list.New()
	lst.PushBack(start)

	// iterate through the possible positions to search for prize
	for lst.Len() > 0 {
		e := lst.Front()
		cur := e.Value.(sPos)
		lst.Remove(e)

		// check for prize
		if cur.x == m.prize.x && cur.y == m.prize.y {
			debugPrintf("  found prize at (%v,%v) with a_steps=%v, b_steps=%v\n", cur.x, cur.y, cur.a_steps, cur.b_steps)
			return cur.a_steps*3 + cur.b_steps
		}

		// check for overshoot (we cannot backtrace)
		if cur.x > m.prize.x || cur.y > m.prize.y {
			continue // skip exploring this way
		}

		// the hint says you should not need to press a or b over 100 times
		if cur.a_steps > 100 || cur.b_steps > 100 {
			debugPrintf("  exceeded 100 steps: a_steps=%v, b_steps=%v\n", cur.a_steps, cur.b_steps)
			continue
		}

		// get next search positions
		nextPos := []sPos{
			{x: cur.x + m.a.x, y: cur.y + m.a.y, a_steps: cur.a_steps + 1, b_steps: cur.b_steps},
			{x: cur.x + m.b.x, y: cur.y + m.b.y, a_steps: cur.a_steps, b_steps: cur.b_steps + 1},
		}

		// add the positions to the list
		for _, p := range nextPos {
			if !visited[[2]int{p.x, p.y}] {
				visited[[2]int{p.x, p.y}] = true
				lst.PushBack(p)
			}
		}
	}
	debugPrintf("  no path to prize found\n")
	return 0 // no path to prize found
}

type sPos struct {
	x, y, a_steps, b_steps int
}

type point struct {
	x, y int
}

type machine struct {
	a, b, prize point
}

func parseInput(input []string) (machines []machine) {
	re := regexp.MustCompile(`(Button A|Button B|Prize): X[+=](\d+), Y[+=](\d+)`)
	var a, b, prize point
	for _, line := range input {
		matches := re.FindStringSubmatch(line)
		if len(matches) != 4 {
			continue
		}
		x, _ := strconv.Atoi(matches[2])
		y, _ := strconv.Atoi(matches[3])
		switch matches[1] {
		case "Button A":
			a = point{x: x, y: y}
		case "Button B":
			b = point{x: x, y: y}
		case "Prize":
			prize = point{x: x, y: y}
			machines = append(machines, machine{a: a, b: b, prize: prize})
		}
	}
	return machines
}
