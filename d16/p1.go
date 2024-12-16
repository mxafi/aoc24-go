package main

import (
	"container/heap"
	"log"
	"math"
)

// reindeer map puzzle
// start on 'S' (facing East/right), end on 'E'
// '#' is a wall, '.' is empty
// moving forward increases score by 1 point
// turning +-90 degrees increases score by 1000 points
// you can only move once per coordinate (aka only turn 90deg once, or move forward once)
// lowest score is the best, return that
// do djikstra...
func solveP1(grid [][]rune) int {
	height := len(grid)
	width := len(grid[0])

	// find start and end positions
	var startX, startY, endX, endY int
	foundStart := false
	foundEnd := false

	for i, row := range grid {
		for j, cell := range row {
			if cell == 'S' {
				startX, startY = i, j
				foundStart = true
			}
			if cell == 'E' {
				endX, endY = i, j
				foundEnd = true
			}
		}
	}

	if !foundStart || !foundEnd {
		log.Fatalln("error: grid must contain both S and E")
	}

	// dist[x][y][d] = minimal cost to reach (x,y) facing direction d
	dist := make([][][]int, height)
	for i := range dist {
		dist[i] = make([][]int, width)
		for j := range dist[i] {
			dist[i][j] = []int{math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt}
		}
	}

	// start facing East (direction = 1)
	dist[startX][startY][1] = 0

	// initialize the priority queue
	pq := &priorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &state{x: startX, y: startY, d: 1, cost: 0})

	for pq.Len() > 0 {
		currentState := heap.Pop(pq).(*state)
		x, y, d, cost := currentState.x, currentState.y, currentState.d, currentState.cost

		// if this is not the best state anymore, skip
		if dist[x][y][d] < cost {
			continue
		}

		// check if we've reached the end tile
		if x == endX && y == endY {
			return cost
		}

		// case 1: move forward if no wall
		dx, dy := directions[d][0], directions[d][1]
		nx, ny := x+dx, y+dy
		if nx >= 0 && nx < height && ny >= 0 && ny < width && grid[nx][ny] != '#' {
			forwardCost := cost + 1
			if forwardCost < dist[nx][ny][d] {
				dist[nx][ny][d] = forwardCost
				heap.Push(pq, &state{x: nx, y: ny, d: d, cost: forwardCost})
			}
		}

		// case 2: turn left (cost=1000)
		ndLeft := (d - 1 + 4) % 4
		leftCost := cost + 1000
		if leftCost < dist[x][y][ndLeft] {
			dist[x][y][ndLeft] = leftCost
			heap.Push(pq, &state{x: x, y: y, d: ndLeft, cost: leftCost})
		}

		// case 3: turn right (cost=1000)
		ndRight := (d + 1) % 4
		rightCost := cost + 1000
		if rightCost < dist[x][y][ndRight] {
			dist[x][y][ndRight] = rightCost
			heap.Push(pq, &state{x: x, y: y, d: ndRight, cost: rightCost})
		}
	}

	log.Fatalln("error: no path found")
	return -1
}

// directions: 0=north, 1=east, 2=south, 3=west
var directions = [4][2]int{
	{-1, 0}, // north
	{0, 1},  // east
	{1, 0},  // south
	{0, -1}, // west
}

// state represents a position and direction in the grid
type state struct {
	x, y, d int // coordinates and direction
	cost    int // cost to reach this state
	index   int // index in the priority queue
}

// priorityQueue implements heap.Interface for managing states by lowest cost
type priorityQueue []*state

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	currentState := x.(*state)
	currentState.index = n
	*pq = append(*pq, currentState)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	currentState := old[n-1]
	currentState.index = -1
	*pq = old[0 : n-1]
	return currentState
}
