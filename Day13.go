package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	Day13()
}

type coordinate struct {
	x int
	y int
}

func Day13() {
	input, _ := ioutil.ReadFile("input13.txt")

	var arr2 = strings.Split(string(input), "\n")
	positions := make(map[coordinate]bool)

	for _, line := range arr2 {
		vals := strings.Split(line, ",")
		c := coordinate{}
		c.x, _ = strconv.Atoi(vals[0])
		c.y, _ = strconv.Atoi(vals[1])
		positions[c] = true
	}

	folds := []string{
		"fold along x=655",
		"fold along y=447",
		"fold along x=327",
		"fold along y=223",
		"fold along x=163",
		"fold along y=111",
		"fold along x=81",
		"fold along y=55",
		"fold along x=40",
		"fold along y=27",
		"fold along y=13",
		"fold along y=6",
	}

	for i, fold := range folds {
		getRidOfFoldAlong := strings.ReplaceAll(fold, "fold along ", "")
		values := strings.Split(getRidOfFoldAlong, "=")
		foldVal, _ := strconv.Atoi(values[1])

		for key := range positions {
			if values[0] == "y" && key.y-foldVal > 0 {
				positions[coordinate{key.x, foldVal - (key.y - foldVal)}] = true
				delete(positions, key)
			} else if values[0] == "x" && key.x-foldVal > 0 {
				positions[coordinate{foldVal - (key.x - foldVal), key.y}] = true
				delete(positions, key)
			}
		}
		if i == 0 {
			// Part 1
			fmt.Println("Part 1 is:\n", len(positions), "\n")
		}
	}

	maxX, maxY := 0, 0

	for key := range positions { // find the max x and y, so we can make the result
		maxX = int(math.Max(float64(key.x), float64(maxX)))
		maxY = int(math.Max(float64(key.y), float64(maxY)))
	}

	arr := [][]string{}

	for i := 0; i < maxY+1; i++ { // making the matrix
		arr = append(arr, make([]string, maxX+1))
	}

	for i := 0; i < len(arr); i++ { // putting in the values of the matrix
		for j := 0; j < len(arr[0]); j++ {
			if positions[coordinate{j, i}] {
				arr[i][j] = "+"
			} else {
				arr[i][j] = " " // putting in this space, so the "+" line up
			}
		}
	}

	fmt.Println("Part 2 is:")

	for _, line := range arr { // Printing out
		fmt.Println(line)
	}
}
