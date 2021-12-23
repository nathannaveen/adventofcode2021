package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(twoDay17())
	fmt.Println(ymin, ymax, xmin, xmax)
}

var ymax, ymin = 0, 0
var xmax, xmin = 0, 0

func parseInput() {
	var input, _ = ioutil.ReadFile("input17.txt")
	s := string(input[:])
	s = s[13:]
	arr := strings.Split(s, " ")
	y := strings.Split(arr[1], "..")
	x := strings.Split(arr[0], "..")
	y[0], x[0] = y[0][2:], x[0][2:]
	x[1] = x[1][:len(x[1])-1]
	ymin, _ = strconv.Atoi(y[0])
	ymax, _ = strconv.Atoi(y[1])
	xmin, _ = strconv.Atoi(x[0])
	xmax, _ = strconv.Atoi(x[1])
}

func oneDay17() int {
	parseInput()
	return ymin * (ymin + 1) / 2
}

func twoDay17() int {
	parseInput()
	res := 0

	for i := ymin; i <= ymin*(ymin+1)/2; i++ {
		for j := 0; j <= xmax; j++ {
			res += canHit(velocity{j, i})
		}
	}

	return res
}

type velocity struct {
	x int
	y int
}

type Pair struct {
	a, b velocity
}

func simulate(pos, v velocity) Pair {
	newPos := velocity{}
	newVelocity := velocity{}

	newPos.x = pos.x + v.x
	newPos.y = pos.y + v.y

	newVelocity.y = v.y - 1
	if v.x > 0 {
		newVelocity.x = v.x - 1
	}
	if v.x < 0 {
		newVelocity.x = v.x + 1
	}
	return Pair{newPos, newVelocity}
}

func canHit(vel velocity) int {
	pos := velocity{0, 0}
	v := vel
	for true {
		if isPast(pos, v) {
			return 0
		}
		if pos.x <= xmax && pos.x >= xmin && pos.y <= ymax && pos.y >= ymin {
			return 1
		}
		temp := simulate(pos, v)
		pos = temp.a
		v = temp.b
	}
	return 1
}

func isPast(pos, v velocity) bool {
	if v.x > 0 && pos.x > xmax {
		return true
	}
	if v.x < 0 && pos.x < xmin {
		return true
	}
	if v.y < 0 && pos.y < ymin {
		return true
	}
	return false
}
