package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func readFileToArray(filename string) []string {
	dat, err := ioutil.ReadFile(filename)
	check(err)

	lines := strings.Split(string(dat), "\n")
	return lines
}

func spacesDirection(movement string) (string, int) {
	getDirection, err := regexp.Compile("[RDUL]{1}")
	check(err)
	getSpaces, err2 := regexp.Compile("\\d+")
	check(err2)
	direction := getDirection.FindString(movement)
	spacesString := getSpaces.FindString(movement)
	if spacesString != "" {
		spaces, err3 := strconv.Atoi(spacesString)

		check(err3)
		return direction, spaces
	}
	return "", 0
}

type Segment struct {
	wire          int
	startX        int
	startY        int
	endX          int
	endY          int
	previousSteps int
}

func getIntersections(segments [][]Segment) [][]Segment {
	intersects := [][]Segment{}

	for _, a := range segments[0] {
		for _, b := range segments[1] {
			if a.endX == b.endX && a.endY == b.endY {
				intersects = append(intersects, []Segment{a, b})
			}
		}
	}
	return intersects
}

func getMinDistance(intersects [][]Segment) float64 {
	minDistance := 999999.0
	oX, oY := 0, 0
	for _, intersect := range intersects {
		distance := math.Abs(float64(intersect[0].endX-oX)) + math.Abs(float64(intersect[0].endY-oY))
		if distance < minDistance && distance != 0 {
			minDistance = distance
		}

	}
	return minDistance
}

func getMinStepsAndDistance(intersects [][]Segment) (int, float64) {
	minDistance, minSteps := 999999.0, 999999
	oX, oY := 0, 0
	for _, intersect := range intersects {
		distance := math.Abs(float64(intersect[0].endX-oX)) + math.Abs(float64(intersect[0].endY-oY))
		steps := intersect[0].previousSteps + intersect[1].previousSteps
		if distance < minDistance && steps < minSteps && distance != 0 {
			minSteps = steps
			minDistance = distance
		}
	}
	return minSteps, minDistance
}

func main() {
	// input := readFileToArray("input.txt")
	input := []string{
		"R75,D30,R83,U83,L12,D49,R71,U7,L72",
		"U62,R66,U55,R34,D71,R55,D58,R83"}
	segments := [][]Segment{[]Segment{}, []Segment{}}
	for a, wire := range input {
		wireMovements := strings.Split(wire, ",")
		x, y, totalSteps := 0, 0, 0
		for _, movement := range wireMovements {
			direction, spaces := spacesDirection(movement)

			for s := 1; s <= spaces; s++ {
				totalSteps++
				if direction == "R" {
					segments[a] = append(segments[a], Segment{wire: a, startX: x, endX: x + 1, startY: y, endY: y, previousSteps: totalSteps})
					x++
				}
				if direction == "L" {
					segments[a] = append(segments[a], Segment{wire: a, startX: x, endX: x - 1, startY: y, endY: y, previousSteps: totalSteps})
					x--
				}
				if direction == "U" {
					segments[a] = append(segments[a], Segment{wire: a, startX: x, endX: x, startY: y, endY: y + 1, previousSteps: totalSteps})
					y++
				}
				if direction == "D" {
					segments[a] = append(segments[a], Segment{wire: a, startX: x, endX: x, startY: y, endY: y - 1, previousSteps: totalSteps})
					y--
				}

			}

		}
	}

	intersects := getIntersections(segments)

	minDistance := getMinDistance(intersects)

	minSteps, minDistance2 := getMinStepsAndDistance(intersects)

	fmt.Println(minDistance, minSteps, minDistance2)
}
