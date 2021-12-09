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

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			isLow := func(y, x int) bool {
				if y >= 0 && x >= 0 && y < len(arr) && x < len(arr[0]) {
					return arr[i][j] < arr[y][x]
				}
				return true
			}
			if isLow(i+1, j) && isLow(i-1, j) && isLow(i, j+1) && isLow(i, j-1) {
				res += int(arr[i][j]-'0') + 1
			}
		}
	}

	fmt.Println(res)
}

// Part Two Bellow:

type lowPos struct {
	actual pos
	start  pos // the low position that we started at
}

type pos struct {
	i, j int
}

func twoDay9() {
	input, _ := ioutil.ReadFile("input9.txt")
	var arr = strings.Split(string(input), "\n")

	stack := []lowPos{}
	m := make(map[pos]int) // position : index in res
	res := []int{}         // size of basin
	visited := make(map[pos]bool)

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			isLow := func(y, x int) bool {
				if y >= 0 && x >= 0 && y < len(arr) && x < len(arr[0]) {
					return arr[i][j] < arr[y][x]
				}
				return true
			}
			if isLow(i+1, j) && isLow(i-1, j) && isLow(i, j+1) && isLow(i, j-1) {
				stack = append(stack, lowPos{pos{i, j}, pos{i, j}}) // find all low values
				res = append(res, 1)
				m[pos{i, j}] = len(res) - 1
			}
		}
	}

	canFlow := func(i, j, val int, start pos) {
		if i >= 0 && j >= 0 && i < len(arr) && j < len(arr[0]) &&
			!visited[pos{i, j}] && arr[i][j] != '9' && int(arr[i][j]-'0') > val {
			visited[pos{i, j}] = true
			res[m[start]]++
			stack = append(stack, lowPos{pos{i, j}, start})
		}
	}

	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		i, j, start := pop.actual.i, pop.actual.j, pop.start

		canFlow(i+1, j, int(arr[i][j]-'0'), start)
		canFlow(i-1, j, int(arr[i][j]-'0'), start)
		canFlow(i, j+1, int(arr[i][j]-'0'), start)
		canFlow(i, j-1, int(arr[i][j]-'0'), start)
	}

	sort.Ints(res)

	fmt.Println(res[len(res)-1] * res[len(res)-2] * res[len(res)-3])
}
