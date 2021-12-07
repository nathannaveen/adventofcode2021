package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type points struct {
	x1, y1 int
	x2, y2 int
}

/*
Made a min and max func because doing:
int(math.Min(float64(a), float64(b))) and int(math.Max(float64(a), float64(b))) is really annoying.
*/

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	oneDay5()
	twoDay5()
}

func oneDay5() {
	file, err := os.Open("input5.txt") // opening the file

	if err != nil {
		panic(err)
	}
	defer file.Close()

	verticals := []points{}
	horizontals := []points{}

	arr := [1002][1002]int{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		temp := scanner.Text() // getting the line
		bigPoints := strings.Split(temp, " -> ")
		points1 := strings.Split(bigPoints[0], ",")
		points2 := strings.Split(bigPoints[1], ",")

		x1, _ := strconv.Atoi(points1[0])
		y1, _ := strconv.Atoi(points1[1])

		x2, _ := strconv.Atoi(points2[0])
		y2, _ := strconv.Atoi(points2[1])

		if y1 == y2 {
			horizontals = append(horizontals, points{min(x1, x2), min(y1, y2), max(x1, x2), max(y1, y2)})
		} else if x1 == x2 {
			verticals = append(verticals, points{min(x1, x2), min(y1, y2), max(x1, x2), max(y1, y2)})
		}
	}

	/*

		Now we can start using line sweep for the horizontal and vertical lines

	*/

	for _, horizontal := range horizontals {
		arr[horizontal.y1][horizontal.x1]++   // start at x1
		arr[horizontal.y2][horizontal.x2+1]-- // remove at x2
	}

	for _, vertical := range verticals {
		for i := vertical.y1; i <= vertical.y2; i++ {
			arr[i][vertical.x1]++
			arr[i][vertical.x1+1]--
		}
	}

	sum := 0
	res := 0

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			sum += arr[i][j]
			if sum > 1 {
				res++
			}
		}
	}

	fmt.Println(res)
}

func twoDay5() {
	file, err := os.Open("input5.txt")
	arr := [1002][1002]int{}
	verticals := []points{}
	horizontals := []points{}
	diagonals := []points{}

	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		temp := scanner.Text()
		bigPoints := strings.Split(temp, " -> ")
		points1 := strings.Split(bigPoints[0], ",")
		points2 := strings.Split(bigPoints[1], ",")

		x1, _ := strconv.Atoi(points1[0])
		y1, _ := strconv.Atoi(points1[1])

		x2, _ := strconv.Atoi(points2[0])
		y2, _ := strconv.Atoi(points2[1])

		if y1 == y2 {
			horizontals = append(horizontals, points{min(x1, x2), min(y1, y2), max(x1, x2), max(y1, y2)})
		} else if x1 == x2 {
			verticals = append(verticals, points{min(x1, x2), min(y1, y2), max(x1, x2), max(y1, y2)})
		} else {
			diagonals = append(diagonals, points{x1, y1, x2, y2})
		}
	}

	for _, horizontal := range horizontals {
		arr[horizontal.y1][horizontal.x1]++
		arr[horizontal.y2][horizontal.x2+1]--
	}

	for _, vertical := range verticals {
		for i := vertical.y1; i <= vertical.y2; i++ {
			arr[i][vertical.x1]++
			arr[i][vertical.x1+1]--
		}
	}

	for _, diagonal := range diagonals {
		y1, y2 := diagonal.y1, diagonal.y2
		x1, x2 := diagonal.x1, diagonal.x2
		if diagonal.y1 > diagonal.y2 {
			y1, y2, x1, x2 = y2, y1, x2, x1
		}

		if x1 < x2 {
			for i, j := y1, x1; i <= y2 && j <= x2; i, j = i+1, j+1 {
				arr[i][j]++
				arr[i][j+1]--
			}
		} else {
			for i, j := y1, x1; i <= y2 && j >= x2; i, j = i+1, j-1 {
				arr[i][j]++
				arr[i][j+1]--
			}
		}
	}

	sum := 0
	res := 0

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			sum += arr[i][j]
			if sum > 1 {
				res++
			}
		}
	}

	fmt.Println(res)
}
