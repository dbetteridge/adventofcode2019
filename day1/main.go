package main

import (
	"fmt"
	"io/ioutil"
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

func calcFuel(mass int, sum int) int {
	fuel := (mass / 3) - 2

	if fuel > 0 {
		sum += fuel
		sum = calcFuel(fuel, sum)
		return sum
	}

	return sum
}

func main() {
	lines := readFileToArray("input.txt")
	sum := 0
	for _, x := range lines {
		line, err := strconv.ParseInt(x, 10, 64)
		if err == nil {
			sum = calcFuel(int(line), sum)
		}
	}
	fmt.Println(sum)
}
