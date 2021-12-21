package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(oneDay17())
}

func calcMinimumY() int {
	var input, _ = ioutil.ReadFile("input17.txt")
	s := string(input[13:])
	y := strings.Split(strings.Fields(s)[1], "..")
	y[0] = y[0][2:] // minimum y, removing "y="
	ymin, _ := strconv.Atoi(y[0])
	return ymin
}

func oneDay17() int {
	ymin := calcMinimumY()
	return ymin * (ymin + 1) / 2
}
