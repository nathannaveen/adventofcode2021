package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	fmt.Println(oneDay10(), twoDay10())
}

func oneDay10() int {
	input, _ := ioutil.ReadFile("input10.txt")
	var arr = strings.Split(string(input), "\n")
	res := 0

	for _, line := range arr {
		stack := []string{}
		shouldBreak := false
		for i := 0; i < len(line); i++ {
			char := string(line[i])
			if char == "(" || char == "[" || char == "{" || char == "<" {
				stack = append(stack, char)
			} else {
				pop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				switch {
				case char == ")" && pop != "(":
					res += 3
					shouldBreak = true
				case char == "]" && pop != "[":
					res += 57
					shouldBreak = true
				case char == "}" && pop != "{":
					res += 1197
					shouldBreak = true
				case char == ">" && pop != "<":
					res += 25137
					shouldBreak = true
				}
				if shouldBreak {
					break
				}
			}
		}
	}

	return res
}

func twoDay10() int {
	input, _ := ioutil.ReadFile("input10.txt")
	var arr = strings.Split(string(input), "\n")
	res := []int{}

	for _, line := range arr {
		stack := []string{}
		shouldBreak := false
		for i := 0; i < len(line); i++ {
			char := string(line[i])
			if char == "(" || char == "[" || char == "{" || char == "<" {
				stack = append(stack, char)
			} else {
				pop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				switch {
				case char == ")" && pop != "(":
					shouldBreak = true
				case char == "]" && pop != "[":
					shouldBreak = true
				case char == "}" && pop != "{":
					shouldBreak = true
				case char == ">" && pop != "<":
					shouldBreak = true
				}
				if shouldBreak {
					break
				}
			}
		}

		if !shouldBreak { // if the line is correct but incomplete
			res = append(res, 0)

			for len(stack) != 0 {
				res[len(res)-1] *= 5
				pop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if pop == "(" {
					res[len(res)-1] += 1
				} else if pop == "[" {
					res[len(res)-1] += 2
				} else if pop == "{" {
					res[len(res)-1] += 3
				} else {
					res[len(res)-1] += 4
				}
			}
		}
	}

	sort.Ints(res)

	return res[len(res)/2]
}
