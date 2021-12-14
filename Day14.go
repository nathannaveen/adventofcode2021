package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	Day14()
}

func Day14() {
	fullInput, _ := ioutil.ReadFile("input14.txt")
	var input = strings.Split(string(fullInput), "\n")

	m := map[string]string{}
	count := map[string]int{}
	s := input[0]

	for i := 2; i < len(input); i++ {
		arr := strings.Split(input[i], " -> ")
		m[arr[0]] = arr[1]
	}

	for i := 2; i <= len(s); i++ {
		count[s[i-2:i]]++
	}

	addIn := func() map[string]int {
		newMap := make(map[string]int)
		for k, v := range count { // copy count
			newMap[k] = v
		}
		for key, val := range count {
			newMap[key] -= val
			newMap[string(key[0])+m[key]] += val
			newMap[m[key]+string(key[1])] += val
		}
		return newMap
	}

	findRes := func() int {
		minimum, maximum := math.MaxInt, 0

		letters := make([]int, 26)

		for key, val := range count {
			arr := []byte(key)
			letters[arr[0]-65] += val
			letters[arr[1]-65] += val
		}

		for _, n := range letters {
			temp := (n + 1) / 2
			if n != 0 {
				if temp < minimum {
					minimum = temp
				}
				if temp > maximum {
					maximum = temp
				}
			}
		}
		return maximum - minimum
	}

	for i := 1; i <= 40; i++ {
		count = addIn()
		if i == 10 {
			fmt.Println("Part 1 day 14 is:", findRes())
		} else if i == 40 {
			fmt.Println("Part 2 day 14 is:", findRes())
		}
	}
}
