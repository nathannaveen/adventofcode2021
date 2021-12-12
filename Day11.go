package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(oneDay11())
}

type point struct {
	i, j int
}

var arr = [][]int{}

func oneDay11() int {
	input, _ := ioutil.ReadFile("input11.txt")

	var arr2 = strings.Split(string(input), "\n")
	arr = make([][]int, len(arr2))

	for i, line := range arr2 {
		arr[i] = make([]int, len(line))
		for j, letter := range line {
			n, _ := strconv.Atoi(string(letter))
			arr[i][j] = n
		}
	}

	flashes := 0
	steps := 0

	for flashes < 100 {
		flashes = 0
		steps++

		for i := 0; i < len(arr); i++ {
			for j := 0; j < len(arr[0]); j++ {
				stack := []point{{i, j}}

				for len(stack) > 0 {
					pop := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					if pop.i >= 0 && pop.j >= 0 && pop.i < len(arr) && pop.j < len(arr[0]) {
						arr[pop.i][pop.j]++
						if arr[pop.i][pop.j] == 10 {
							flashes++

							stack = append(stack, point{pop.i + 1, pop.j}, point{pop.i - 1, pop.j}, point{pop.i, pop.j + 1}, point{pop.i, pop.j - 1},
								point{pop.i + 1, pop.j + 1}, point{pop.i - 1, pop.j - 1}, point{pop.i + 1, pop.j - 1}, point{pop.i - 1, pop.j + 1})
							//if pop.i - 1 >= 0 {
							//	stack = append(stack, point{pop.i - 1, pop.j})
							//}
							//if pop.i - 1 >= 0 && pop.j + 1< len(arr[0]) {
							//	stack = append(stack, point{pop.i - 1, pop.j + 1})
							//}
							//if pop.j + 1 < len(arr[0]) {
							//	stack = append(stack, point{i, pop.j + 1})
							//}
							//if pop.i + 1 < len(arr) && pop.j + 1 < len(arr[0]) {
							//	stack = append(stack, point{pop.i + 1, pop.j + 1})
							//}
							//if pop.i + 1 < len(arr) {
							//	stack = append(stack, point{pop.i + 1, pop.j})
							//}
							//if pop.i < len(arr) && pop.j - 1 >= 0 {
							//	stack = append(stack, point{pop.i + 1, pop.j - 1})
							//}
							//if pop.j - 1 >= 0 {
							//	stack = append(stack, point{i, pop.j - 1})
							//}
							//if pop.i - 1 >= 0 && pop.j - 1 >= 0 {
							//	stack = append(stack, point{pop.i - 1, pop.j - 1})
							//}
						}
					}
				}
			}
		}

		for i := 0; i < len(arr); i++ {
			for j := 0; j < len(arr[0]); j++ {
				if arr[i][j] > 9 {
					arr[i][j] = 0
				}
			}
		}
	}

	return steps
}
