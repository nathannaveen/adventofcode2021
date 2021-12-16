package main

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("Part one day 15 is:", oneDay15())
	fmt.Println("Part two day 15 is:", twoDay15())
}

type pos struct {
	i, j int
	risk int
}

type PosHeap []pos // minHeap of our struct pos

func (h PosHeap) Len() int            { return len(h) }
func (h PosHeap) Less(i, j int) bool  { return h[i].risk < h[j].risk }
func (h PosHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *PosHeap) Push(x interface{}) { *h = append(*h, x.(pos)) }

func (h *PosHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

var arr = [][]int{}

func getInput() {
	fullInput, _ := ioutil.ReadFile("input15.txt")
	var input = strings.Split(string(fullInput), "\n")

	for _, line := range input { // make arr from the input
		arr = append(arr, make([]int, len(line)))
		for i := 0; i < len(line); i++ {
			arr[len(arr)-1][i] = int(line[i] - '0')
		}
	}
}

func oneDay15() int {
	arr = [][]int{}
	getInput()
	visited := map[pos]bool{}

	h := &PosHeap{pos{0, 0, 0}}
	heap.Init(h)

	n := len(arr)

	move := func(i, j, risk int) {
		if i >= 0 && j >= 0 && i < n && j < n && !visited[pos{i, j, 0}] {
			heap.Push(h, pos{i, j, risk + arr[i][j]})
			visited[pos{i, j, 0}] = true
		}
	}

	for h.Len() != 0 {
		pop := heap.Pop(h).(pos)

		if pop.i == n-1 && pop.j == n-1 {
			return pop.risk
		}

		move(pop.i+1, pop.j, pop.risk) // Doing the 4 adjacent positions
		move(pop.i-1, pop.j, pop.risk)
		move(pop.i, pop.j+1, pop.risk)
		move(pop.i, pop.j-1, pop.risk)
	}

	return -1 // there is no input
}

func twoDay15() int {
	arr = [][]int{}
	getInput()
	visited := map[pos]bool{}
	h := &PosHeap{pos{0, 0, 0}}
	heap.Init(h)

	n := len(arr)
	totalN := n * 5

	move := func(i, j, risk int) {
		if i >= 0 && j >= 0 && i < totalN && j < totalN && !visited[pos{i, j, 0}] {
			temp := (arr[i%n][j%n]+i/n+j/n-1)%9 + 1 // main change between P1 and P2 is this line
			heap.Push(h, pos{i, j, risk + temp})
			visited[pos{i, j, 0}] = true
		}
	}

	for h.Len() != 0 {
		pop := heap.Pop(h).(pos)

		if pop.i == totalN-1 && pop.j == totalN-1 {
			return pop.risk
		}

		move(pop.i+1, pop.j, pop.risk)
		move(pop.i-1, pop.j, pop.risk)
		move(pop.i, pop.j+1, pop.risk)
		move(pop.i, pop.j-1, pop.risk)
	}

	return -1
}
