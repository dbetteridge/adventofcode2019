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

func stacker(layers [][][]int) [6][25]int {
	result := [6][25]int{}
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			for k := len(layers) - 1; k >= 0; k-- {
				newV := layers[k][i][j]
				if newV != 2 {
					result[i][j] = newV
				}

			}
		}
	}
	return result
}

func main() {
	w, h := 25, 6
	lines := readFileToArray("./input.txt")
	line := lines[0]
	layers := [][][]int{}
	for i := 0; i < len(line); i += w * h {
		layer := [][]int{}
		row := []int{}
		for j := 0; j < w*h; j++ {
			value, err := strconv.Atoi(string(line[i+j]))
			check(err)
			row = append(row, value)
			if len(row) == w {
				layer = append(layer, row)
				row = []int{}
			}
		}
		layers = append(layers, layer)
	}
	result := stacker(layers)
	for _, row := range result {
		for _, v := range row {
			if v == 0 {
				fmt.Printf(" ")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}
