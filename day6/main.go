package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFileToArray(filename string) []string {
	dat, err := ioutil.ReadFile(filename)
	check(err)

	lines := strings.Split(string(dat), "\n")
	return lines
}

type Orbit struct {
	name  string
	child string
}

func rec(orbits map[string]string, key string, totalOrbits int) int {
	if orbits[key] != "COM" {
		totalOrbits++
		return rec(orbits, orbits[key], totalOrbits)
	}
	return totalOrbits
}

func getTotalOrbits(orbits map[string]string, orbitStrs []string, totalOrbits int) {
	for _, orbitStr := range orbitStrs {
		orbitArr := strings.Split(orbitStr, ")")
		if len(orbitArr) < 2 {
			break
		}
		parent := orbitArr[0]
		key := orbitArr[1]
		orbits[key] = parent
		totalOrbits++
	}
	for key := range orbits {
		totalOrbits = rec(orbits, key, totalOrbits)
	}

	fmt.Println(totalOrbits)
}

func visitedOrbits(orbits map[string]string, key string, visited []string) []string {
	if orbits[key] != "COM" {
		visited = append(visited, orbits[key])
		return visitedOrbits(orbits, orbits[key], visited)
	}
	return visited
}

func findCommon(visitedYOU []string, visitedSANTA []string) (int, int) {
	for y, visitedY := range visitedYOU {
		for s, visitedS := range visitedSANTA {
			if visitedY == visitedS {
				return y, s
			}
		}
	}
	return 0, 0
}

func main() {
	orbitStrs := readFileToArray("input.txt")
	var orbits map[string]string
	orbits = make(map[string]string)
	totalOrbits := 0

	for _, orbitStr := range orbitStrs {
		orbitArr := strings.Split(orbitStr, ")")
		if len(orbitArr) < 2 {
			break
		}
		parent := orbitArr[0]
		key := orbitArr[1]
		orbits[key] = parent
	}
	visitedYOU := []string{}
	visitedSANTA := []string{}

	visitedYOU = visitedOrbits(orbits, "YOU", visitedYOU)
	visitedSANTA = visitedOrbits(orbits, "SAN", visitedSANTA)
	countY, countS := findCommon(visitedYOU, visitedSANTA)
	totalOrbits = countS + countY
	fmt.Println(totalOrbits)
}
