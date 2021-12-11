package main

import "fmt"

func main() {
	oneDay6()
	twoDay6()
}

func oneDay6() {
	arr := []int{
		4, 1, 1, 4, 1, 2, 1, 4, 1, 3, 4, 4, 1, 5, 5, 1, 3, 1, 1, 1, 4, 4, 3, 1, 5, 3, 1, 2, 5, 1, 1, 5, 1, 1, 4, 1, 1,
		1, 1, 2, 1, 5, 3, 4, 4, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 5, 1, 1, 1, 4, 1, 2, 3, 5, 1, 2, 2, 4, 1,
		4, 4, 4, 1, 2, 5, 1, 2, 1, 1, 1, 1, 1, 1, 4, 1, 1, 4, 3, 4, 2, 1, 3, 1, 1, 1, 3, 5, 5, 4, 3, 4, 1, 5, 1, 1, 1,
		2, 2, 1, 3, 1, 2, 4, 1, 1, 3, 3, 1, 3, 3, 1, 1, 3, 1, 5, 1, 1, 3, 1, 1, 1, 5, 4, 1, 1, 1, 1, 4, 1, 1, 3, 5, 4,
		3, 1, 1, 5, 4, 1, 1, 2, 5, 4, 2, 1, 4, 1, 1, 1, 1, 3, 1, 1, 1, 1, 4, 1, 1, 1, 1, 2, 4, 1, 1, 1, 1, 3, 1, 1, 5,
		1, 1, 1, 1, 1, 1, 4, 2, 1, 3, 1, 1, 1, 2, 4, 2, 3, 1, 4, 1, 2, 1, 4, 2, 1, 4, 4, 1, 5, 1, 1, 4, 4, 1, 2, 2, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 4, 5, 4, 1, 3, 1, 3, 1, 1, 1, 5, 3, 5, 5, 2, 2, 1, 4, 1, 4, 2, 1, 4, 1, 2, 1, 1,
		2, 1, 1, 5, 4, 2, 1, 1, 1, 2, 4, 1, 1, 1, 1, 2, 1, 1, 5, 1, 1, 2, 2, 5, 1, 1, 1, 1, 1, 2, 4, 2, 3, 1, 2, 1, 5,
		4, 5, 1, 4}

	for j := 0; j < 80; j++ {
		for i := 0; i < len(arr); i++ {
			if arr[i] == 0 {
				arr[i] = 7
				arr = append(arr, 9)
			}
			arr[i]--
		}

	}
	fmt.Println(len(arr))
}

func twoDay6() {
	arr := []int{
		4, 1, 1, 4, 1, 2, 1, 4, 1, 3, 4, 4, 1, 5, 5, 1, 3, 1, 1, 1, 4, 4, 3, 1, 5, 3, 1, 2, 5, 1, 1, 5, 1, 1, 4, 1, 1,
		1, 1, 2, 1, 5, 3, 4, 4, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 5, 1, 1, 1, 4, 1, 2, 3, 5, 1, 2, 2, 4, 1,
		4, 4, 4, 1, 2, 5, 1, 2, 1, 1, 1, 1, 1, 1, 4, 1, 1, 4, 3, 4, 2, 1, 3, 1, 1, 1, 3, 5, 5, 4, 3, 4, 1, 5, 1, 1, 1,
		2, 2, 1, 3, 1, 2, 4, 1, 1, 3, 3, 1, 3, 3, 1, 1, 3, 1, 5, 1, 1, 3, 1, 1, 1, 5, 4, 1, 1, 1, 1, 4, 1, 1, 3, 5, 4,
		3, 1, 1, 5, 4, 1, 1, 2, 5, 4, 2, 1, 4, 1, 1, 1, 1, 3, 1, 1, 1, 1, 4, 1, 1, 1, 1, 2, 4, 1, 1, 1, 1, 3, 1, 1, 5,
		1, 1, 1, 1, 1, 1, 4, 2, 1, 3, 1, 1, 1, 2, 4, 2, 3, 1, 4, 1, 2, 1, 4, 2, 1, 4, 4, 1, 5, 1, 1, 4, 4, 1, 2, 2, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 4, 5, 4, 1, 3, 1, 3, 1, 1, 1, 5, 3, 5, 5, 2, 2, 1, 4, 1, 4, 2, 1, 4, 1, 2, 1, 1,
		2, 1, 1, 5, 4, 2, 1, 1, 1, 2, 4, 1, 1, 1, 1, 2, 1, 1, 5, 1, 1, 2, 2, 5, 1, 1, 1, 1, 1, 2, 4, 2, 3, 1, 2, 1, 5,
		4, 5, 1, 4}

	temp := make([]int, 9)

	for _, a := range arr {
		temp[a]++
	}

	for j := 0; j < 256; j++ {
		zero := temp[0]
		temp[0] = 0
		for i := 1; i <= 8; i++ {
			temp[i-1] = temp[i]
		}
		temp[8] = zero
		temp[6] += zero
	}

	res := 0

	for _, t := range temp {
		res += t
	}

	fmt.Println(res)
}