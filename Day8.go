package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func main() {
	oneDay8()
	twoDay8()
}

func oneDay8() {
	file, err := os.Open("input8.txt") // opening the file

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	res := 0

	for scanner.Scan() {
		temp := scanner.Text() // getting the line
		arr := strings.Split(temp, " | ")
		words := strings.Split(arr[1], " ")

		for _, s := range words {
			if len(s) == 2 || len(s) == 4 || len(s) == 3 || len(s) == 7 {
				res++
			}
		}
	}
	fmt.Println(res)
}

/*
Part two is a lot harder than part one, so I wrote an explanation about it.
*/

func twoDay8() {
	b, _ := ioutil.ReadFile("input8.txt")
	var h = strings.Split(string(b), "\n")

	res := 0

	for _, temp := range h {
		after := strings.Split(strings.Split(temp, " | ")[1], " ")
		before := strings.Split(strings.Split(temp, " | ")[0], " ")
		jumbledNumber := make(map[int]string)
		properNumber := make(map[string]int)
		// jumbledNumber = number : mixed up numbers (used to find the number in O(1) time)
		// properNumber = mixed up number : number (used for finding the mixed up number in O(1) time)
		num := 0
		arr := []string{}

		lengthFiveOrSix := func(a string) {
			fourCounter, oneCounter := 0, 0
			four, one := jumbledNumber[4], jumbledNumber[1]
			newA := strings.Split(a, "")
			sort.Strings(newA)
			a = strings.Join(newA, "")

			for i := 0; i < len(a); i++ {
				if !strings.Contains(four, string(a[i])) {
					fourCounter++
				}
				if !strings.Contains(one, string(a[i])) {
					oneCounter++
				}
			}

			if len(a) == 5 {
				if fourCounter == 3 {
					jumbledNumber[2] = a
					properNumber[a] = 2
				} else if oneCounter == 3 {
					jumbledNumber[3] = a
					properNumber[a] = 3
				} else {
					jumbledNumber[5] = a
					properNumber[a] = 5
				}
			} else { // len(a) == 6
				if fourCounter == 2 {
					jumbledNumber[9] = a
					properNumber[a] = 9
				} else if oneCounter == 5 {
					jumbledNumber[6] = a
					properNumber[a] = 6
				} else {
					jumbledNumber[0] = a
					properNumber[a] = 0
				}
			}
		}

		for _, a := range before {
			newA := strings.Split(a, "")
			sort.Strings(newA)
			a = strings.Join(newA, "")
			switch len(a) {
			case 2:
				jumbledNumber[1] = a
				properNumber[a] = 1
			case 3:
				jumbledNumber[7] = a
				properNumber[a] = 7
			case 4:
				jumbledNumber[4] = a
				properNumber[a] = 4
			case 7:
				jumbledNumber[8] = a
				properNumber[a] = 8
			default: // if the length is 5 or 6
				if jumbledNumber[4] != "" && jumbledNumber[1] != "" {
					lengthFiveOrSix(a)
				} else {
					// we need the 1 and 4 to find out the actual number,
					// so skip for now
					arr = append(arr, a)
				}
			}
		}

		for _, a := range arr { // Do all the skipped number
			lengthFiveOrSix(a)
		}

		for _, a := range after {
			newA := strings.Split(a, "")
			sort.Strings(newA)
			a = strings.Join(newA, "")

			num = (num * 10) + properNumber[a]
		}
		res += num
	}

	fmt.Println(res)
}
