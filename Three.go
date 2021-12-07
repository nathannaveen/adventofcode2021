package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	two()
}

func two() {
	file, err := os.Open("input3.txt")

	if err != nil {
		panic(err)

	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	arr := make([]int, 5)
	arr2 := []string{}
	pos := 1
	one, two := make(map[string]bool), make(map[string]bool)

	for scanner.Scan() {
		temp := scanner.Text()
		arr2 = append(arr2, temp)
		for i := 0; i < len(temp); i++ {
			if temp[i] == '1' {
				arr[i]++
			} else {
				arr[i]--
			}
		}
		one[temp] = true
		two[temp] = true
	}

	for i := 0; i < 5; i++ {
		for _, temp := range arr2 { // not scanning
			if arr[i] > 0 && temp[i] == '0' && len(one) != 1 {
				delete(one, temp)
			} else if arr[i] < 0 && temp[i] == '1' && len(one) != 1 {
				delete(one, temp)
			}

			if arr[i] > 0 && temp[i] == '1' && len(two) != 1 {
				delete(two, temp)
			} else if arr[i] < 0 && temp[i] == '0' && len(two) != 1 {
				delete(two, temp)
			}

		}
	}
	res1 := 0
	res2 := 0
	pos = 1

	for a := range one {
		fmt.Println(a)
		for i := len(a) - 1; i >= 0; i-- {
			if a[i] == '1' {
				res1 += pos
			}
			pos *= 2
		}
	}

	pos = 1
	for a := range two {
		for i := len(a) - 1; i >= 0; i-- {
			if a[i] == '1' {
				res2 += pos
			}
			pos *= 2
		}
	}

	fmt.Println(one, two, res1, res2, res1*res2)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func one() {
	file, err := os.Open("input3.txt")

	if err != nil {
		panic(err)

	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	arr := make([]int, 12)
	gamma := 0
	epsilon := 0
	pos := 1

	for scanner.Scan() {
		temp := scanner.Text()
		for i := 0; i < len(temp); i++ {
			if temp[i] == '1' {
				arr[i]++
			} else {
				arr[i]--
			}
		}
	}

	fmt.Println(arr)

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] > 0 {
			gamma += pos
		} else {
			epsilon += pos
		}
		pos *= 2

	}

	fmt.Println(gamma*epsilon, gamma, epsilon)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
