package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	Day12()
}

var res = 0

func Day12() {
	input, _ := ioutil.ReadFile("input12.txt")

	var arr2 = strings.Split(string(input), "\n")
	connections := map[string][]string{}

	for _, line := range arr2 { //  get everything from input
		s := strings.Split(line, "-")
		connections[s[0]] = append(connections[s[0]], s[1])
		connections[s[1]] = append(connections[s[1]], s[0])
	}

	res = 0
	one(make(map[string]bool), connections, "start")
	fmt.Println("Part ones answer:", res)

	res = 0
	two(make(map[string]bool), connections, "start", false)
	fmt.Println("Part twos answer:", res)
}

func one(visited map[string]bool, connections map[string][]string, cur string) {
	for _, connect := range connections[cur] {
		if connect == "end" {
			res++
		} else if connect == "start" {
			continue
		} else if !visited[connect] {
			newMap := make(map[string]bool)
			for k, v := range visited { // remake visited
				newMap[k] = v
			}
			if connect[0] >= 97 { // if lowercase add to visited
				newMap[connect] = true
			}
			one(newMap, connections, connect)
		}
	}
}

func two(visited map[string]bool, connections map[string][]string, cur string, hasDoneTwice bool) {
	for _, connect := range connections[cur] {
		if connect == "end" {
			res++
			continue
		} else if connect == "start" {
			continue
		}

		newMap := make(map[string]bool)
		for k, v := range visited { // remake visited
			newMap[k] = v
		}
		if connect[0] >= 97 { // if lowercase add to visited
			newMap[connect] = true
		}

		if (!visited[connect] && hasDoneTwice) || (visited[connect] && !hasDoneTwice) {
			two(newMap, connections, connect, true)
		} else if !visited[connect] {
			two(newMap, connections, connect, false)
		}
	}
}
