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

func instructionsFromArray(array []string) []int {
	instructions := []int{}
	for _, x := range array {
		singleInt64, err := strconv.ParseInt(x, 10, 0)
		check(err)
		singleInt := int(singleInt64)
		instructions = append(instructions, singleInt)
	}
	return instructions
}

func runInstructions(inputInstructions []int) int {
	index := 0
	memory := make([]int, len(inputInstructions))
	copy(memory, inputInstructions)
	for index < len(memory)-1 {
		opCode := memory[index]
		firstAddress := index + 1
		secondAddress := firstAddress + 1
		storageAddress := secondAddress + 1
		switch opCode {
		case 99:
			index += 4
		case 1:
			leftValue := memory[memory[firstAddress]]
			rightValue := memory[memory[secondAddress]]
			storeLocation := memory[storageAddress]

			memory[storeLocation] = leftValue + rightValue
			index += 4
		case 2:
			leftValue := memory[memory[firstAddress]]
			rightValue := memory[memory[secondAddress]]
			storeLocation := memory[storageAddress]

			memory[storeLocation] = leftValue * rightValue
			index += 4
		}
	}
	return memory[0]
}

func findNounAndVerb(memory []int) (int, int) {
	noun, verb := 0, 0
	for noun = 0; noun < 99; noun++ {
		memory[1] = noun
		for verb = 0; verb < 99; verb++ {
			memory[2] = verb
			result := runInstructions(memory)
			if result == 19690720 {
				return noun, verb
			}
		}
	}
	return noun, verb
}

func main() {
	lines := readFileToArray("input.txt")
	linesAsArray := strings.Split(lines[0], ",")

	instructions := instructionsFromArray(linesAsArray)

	noun, verb := findNounAndVerb(instructions)
	fmt.Println(noun, verb)
}
