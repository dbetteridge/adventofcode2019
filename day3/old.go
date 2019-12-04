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
		panic(e)
	}
}

func readFileToArray(filename string) []string {
	dat, err := ioutil.ReadFile(filename)
	check(err)

	lines := strings.Split(string(dat), "\n")
	return lines
}

func main() {
	input := readFileToArray("input.txt")
	// input := []string{"R75,D30,R83,U83,L12,D49,R71,U7,L72",
	// 	"U62,R66,U55,R34,D71,R55,D58,R83"}
	getDirection, err := regexp.Compile("[RDUL]{1}")
	getSpaces, err2 := regexp.Compile("\\d+")
	check(err2)
	check(err)
	wireGrid := [50000][50000][2]int{}
	intersects := [][]int{}
	for a, wire := range input {
		wireMovements := strings.Split(wire, ",")
		x, y, totalSteps := 25000, 25000, 0
		for _, movement := range wireMovements {

			direction := getDirection.FindString(movement)
			if len(direction) > 0 {
				spaces, err3 := strconv.Atoi(getSpaces.FindString(movement))
				check(err3)
				switch direction {
				case "R":
					for i := x; i < x+spaces; i++ {
						if wireGrid[i][y][0] == 0 {
							wireGrid[i][y] = [2]int{a + 1, totalSteps + (i - x)}
						} else if wireGrid[i][y][0] == a+1 {
							continue
						} else {
							intersects = append(intersects, []int{i, y, wireGrid[i][y][1], totalSteps + (i - x)})
							wireGrid[i][y] = [2]int{5, totalSteps + (i - x)}
						}
					}
					x += spaces
					totalSteps += spaces
				case "L":
					for i := x; i > x-spaces; i-- {
						if wireGrid[i][y][0] == 0 {
							wireGrid[i][y] = [2]int{a + 1, totalSteps + (x - i)}
						} else if wireGrid[i][y][0] == a+1 {
							continue
						} else {
							intersects = append(intersects, []int{i, y, wireGrid[i][y][1], totalSteps + (x - i)})
							wireGrid[i][y] = [2]int{5, totalSteps + (x - i)}
						}
					}
					x -= spaces
					totalSteps += spaces
				case "U":
					for i := y; i < y+spaces; i++ {
						if wireGrid[x][i][0] == 0 {
							wireGrid[x][i] = [2]int{a + 1, totalSteps + (i - y)}
						} else if wireGrid[x][i][0] == a+1 {
							continue
						} else {
							intersects = append(intersects, []int{x, i, wireGrid[x][i][1], totalSteps + (i - y)})
							wireGrid[x][i] = [2]int{5, totalSteps + (i - y)}
						}
					}
					y += spaces
					totalSteps += spaces
				case "D":
					for i := y; i > y-spaces; i-- {
						if wireGrid[x][i][0] == 0 {
							wireGrid[x][i] = [2]int{a + 1, totalSteps + (y - i)}
						} else if wireGrid[x][i][0] == a+1 {
							continue
						} else {
							intersects = append(intersects, []int{x, i, wireGrid[x][i][1], totalSteps + (y - i)})
							wireGrid[x][i] = [2]int{5, totalSteps + (y - i)}
						}
					}
					y -= spaces
					totalSteps += spaces
				default:
					continue
				}
			}

		}

	}
	iX, iY := 25000, 25000
	minDistance, minSteps := 9999999.0, 9999999
	for _, intersect := range intersects {
		distance := math.Abs(float64(intersect[0]-iX)) + math.Abs(float64(intersect[1]-iY))
		steps := intersect[2] + intersect[3]
		if distance < minDistance && (intersect[0] != iX && intersect[1] != iY) && (steps < minSteps) {
			minDistance = distance
			minSteps = steps
		}
	}

	fmt.Println(minDistance, minSteps)

}
