package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	parseInput()
	fmt.Println(oneDay17())
	fmt.Println(twoDay17())
}

var ymax, ymin = 0, 0
var xmax, xmin = 0, 0

func parseInput() {
	var input, _ = ioutil.ReadFile("input17.txt")
	s := string(input[13:])      // remove "target area: "
	arr := strings.Split(s, " ") // splitting into x values, and y values (ie: ["x=000..000,", "y=000..000"])

	arr[0], arr[1] = arr[0][2:], arr[1][2:] // remove the "x=", and the "y="

	y := strings.Split(arr[1], "..") // split into two y values
	x := strings.Split(arr[0], "..") // split into two x values

	x[1] = x[1][:len(x[1])-1] // remove the "," after the second x value

	ymin, _ = strconv.Atoi(y[0])
	ymax, _ = strconv.Atoi(y[1])
	xmin, _ = strconv.Atoi(x[0])
	xmax, _ = strconv.Atoi(x[1])
}

// PART 1:

func oneDay17() int {
	return ymin * (ymin + 1) / 2
}

// PART 2:

type velocity struct {
	x, y int
}

func twoDay17() int {
	res := 0

	for i := ymin; i <= ymin*(ymin+1)/2; i++ {
		for j := 0; j <= xmax; j++ {
			res += canHit(velocity{j, i})
		}
	}

	return res
}

func canHit(v velocity) int {
	pos := v
	x := v.x
	y := v.y

	for true {
		if pos.x > xmax || pos.y < ymin { // if the probe has passed the target
			return 0
		}
		if pos.x <= xmax && pos.x >= xmin && pos.y <= ymax && pos.y >= ymin { // if the probe is inside the target
			return 1
		}

		// simulating moving/falling probe bellow

		y -= 1
		if x > 0 {
			x -= 1
		} else if x < 0 {
			x += 1
		}

		pos.x += x
		pos.y += y
	}

	return -1 // will never reach this
}
