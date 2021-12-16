package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	oneDay9()
	twoDay9()
}

func oneDay9() {
	input, _ := ioutil.ReadFile("input9.txt")
	var arr = strings.Split(string(input), "\n")

	res := 0
	i, j := 0, 0

	isLow := func(y, x int) bool {
		// I put this function inside twoDay9()
		// because then I wouldn't have to parse i, j, and arr which would make it really messy.
		if y >= 0 && x >= 0 && y < len(arr) && x < len(arr[0]) {
			return arr[i][j] < arr[y][x]
		}
		return true
	}

	for i = 0; i < len(arr); i++ {
		for j = 0; j < len(arr[0]); j++ {
			if isLow(i+1, j) && isLow(i-1, j) && isLow(i, j+1) && isLow(i, j-1) {
				res += int(arr[i][j]-'0') + 1
			}
		}
	}

	fmt.Println(res)
}

// Part Two Bellow:

type lowPos struct {
	actual actualPos
	start  actualPos // the low position that we started at
}

type actualPos struct {
	i, j int
}

func twoDay9() {
	input, _ := ioutil.ReadFile("input9.txt")
	var arr = strings.Split(string(input), "\n")

	stack := []lowPos{}
	m := make(map[actualPos]int) // position : index in res
	res := []int{}               // size of basin
	visited := make(map[actualPos]bool)
	i, j := 0, 0

	isLow := func(y, x int) bool {
		// I put this function inside twoDay9()
		// because then I wouldn't have to parse i, j, and arr which would make it really messy.
		if y >= 0 && x >= 0 && y < len(arr) && x < len(arr[0]) {
			return arr[i][j] < arr[y][x]
		}
		return true
	}

	for i = 0; i < len(arr); i++ {
		for j = 0; j < len(arr[0]); j++ {
			if isLow(i+1, j) && isLow(i-1, j) && isLow(i, j+1) && isLow(i, j-1) {
				stack = append(stack, lowPos{actualPos{i, j}, actualPos{i, j}})
				res = append(res, 1)
				m[actualPos{i, j}] = len(res) - 1
			}
		}
	}

	canFlow := func(i, j, val int, start actualPos) {
		// This is for finding whether smoke can flow from one place to another.
		if i >= 0 && j >= 0 && i < len(arr) && j < len(arr[0]) &&
			!visited[actualPos{i, j}] && arr[i][j] != '9' && int(arr[i][j]-'0') > val {
			visited[actualPos{i, j}] = true
			res[m[start]]++
			stack = append(stack, lowPos{actualPos{i, j}, start})
		}
	}

	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		y, x, start := pop.actual.i, pop.actual.j, pop.start

		canFlow(y+1, x, int(arr[i][j]-'0'), start)
		canFlow(y-1, x, int(arr[i][j]-'0'), start)
		canFlow(y, x+1, int(arr[i][j]-'0'), start)
		canFlow(y, x-1, int(arr[i][j]-'0'), start)
	}

	sort.Ints(res)

	fmt.Println(res[len(res)-1] * res[len(res)-2] * res[len(res)-3]) // Last
}
