package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	oneDay4()
}

type key struct {
	m map[int]point
}

type point struct {
	i int
	j int
}

type key2 struct {
	col map[int][]int
	row map[int][]int
}

func oneDay4() {
	arr := []int{46, 79, 77, 45, 57, 34, 44, 13, 32, 88, 86, 82, 91, 97, 89, 1, 48, 31, 18, 10, 55, 74, 24, 11, 80, 78, 28, 37, 47, 17, 21, 61, 26, 85,
		99, 96, 23, 70, 3, 54, 5, 41, 50, 63, 14, 64, 42, 36, 95, 52, 76, 68, 29, 9, 98, 35, 84, 83, 71, 49, 73, 58, 56, 66, 92, 30, 51, 20, 81, 69, 65, 15,
		6, 16, 39, 43, 67, 7, 59, 40, 60, 4, 90, 72, 22, 0, 93, 94, 38, 53, 87, 27, 12, 2, 25, 19, 8, 62, 33, 75}

	file, err := os.Open("input4.txt")

	if err != nil {
		panic(err)

	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0
	points := []key{}
	m := make(map[int]bool)

	for scanner.Scan() {
		temp := scanner.Text()
		s := strings.Fields(temp)

		if len(s) == 0 {
			i = 0
			points = append(points, key{make(map[int]point)})
		} else {
			for j := 0; j < len(s); j++ {
				b, _ := strconv.Atoi(s[j])
				points[len(points)-1].m[b] = point{i, j}
			}

			i++
		}
	}

	arr2 := []key2{}

	for j := 0; j < len(points); j++ {
		arr2 = append(arr2, key2{make(map[int][]int), make(map[int][]int)})
	}

	for t, a := range arr {
		for j, point2 := range points {
			if _, ok := point2.m[a]; ok {
				if m[j] == false {
					arr2[j].col[point2.m[a].j] = append(arr2[j].col[point2.m[a].j], a)
					arr2[j].row[point2.m[a].i] = append(arr2[j].row[point2.m[a].i], a)

					if len(arr2[j].row[point2.m[a].i]) == 5 || len(arr2[j].col[point2.m[a].j]) == 5 {
						sum := 0

						for x := range point2.m {
							sum += x
						}

						for k := 0; k <= t; k++ {
							if _, ok2 := point2.m[arr[k]]; ok2 {
								sum -= arr[k]
							}
						}

						fmt.Println(sum * a)
						m[j] = true
					}
				}
			}

		}
	}
}
