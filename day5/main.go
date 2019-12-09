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

func getAddressOrValue(memory []int, index int, mode int) int {
	if mode == 0 {
		return memory[memory[index]]
	}
	return memory[index]
}

func setAddressOrValue(memory []int, index int, value int) {

	memory[memory[index]] = value
}

func runInstructions(inputInstructions []int) int {
	index := 0
	memory := make([]int, len(inputInstructions))
	copy(memory, inputInstructions)
	input := 5
	for index < len(memory)-1 {
		opCode := memory[index]
		firstAddress := index + 1
		secondAddress := firstAddress + 1
		storageAddress := secondAddress + 1
		modes := []int{0, 0, 0}
		if opCode > 10 {
			imOpCode := strconv.Itoa(opCode)
			opCode = 0
			modeCount := 5 - len(imOpCode)
			for i, c := range imOpCode {
				intC, err := strconv.Atoi(string(c))
				check(err)
				if i == len(imOpCode)-1 || i == len(imOpCode)-2 {
					opCode += intC
				} else {
					modes[modeCount] = intC
					modeCount++
				}

			}
			fmt.Println(imOpCode, opCode, modes)
		}
		switch opCode {
		case 99:
			index = len(memory) - 1
		case 1:
			leftValue := getAddressOrValue(memory, firstAddress, modes[len(modes)-1])
			rightValue := getAddressOrValue(memory, secondAddress, modes[len(modes)-2])
			setAddressOrValue(memory, storageAddress, leftValue+rightValue)
			index += 4
		case 2:
			leftValue := getAddressOrValue(memory, firstAddress, modes[len(modes)-1])
			rightValue := getAddressOrValue(memory, secondAddress, modes[len(modes)-2])
			setAddressOrValue(memory, storageAddress, leftValue*rightValue)
			index += 4
		case 3:
			setAddressOrValue(memory, index+1, input)
			index += 2
		case 4:
			fmt.Println(getAddressOrValue(memory, index+1, modes[2]))
			index += 2
		case 5:
			leftValue := getAddressOrValue(memory, index+1, modes[len(modes)-1])
			rightValue := getAddressOrValue(memory, index+2, modes[len(modes)-2])
			if leftValue != 0 {
				index = rightValue
			} else {
				index += 3
			}
		case 6:
			leftValue := getAddressOrValue(memory, index+1, modes[len(modes)-1])
			rightValue := getAddressOrValue(memory, index+2, modes[len(modes)-2])
			if leftValue == 0 {
				index = rightValue
			} else {
				index += 3
			}
		case 7:
			leftValue := getAddressOrValue(memory, index+1, modes[len(modes)-1])
			rightValue := getAddressOrValue(memory, index+2, modes[len(modes)-2])
			if leftValue < rightValue {
				setAddressOrValue(memory, storageAddress, 1)
			} else {
				setAddressOrValue(memory, storageAddress, 0)
			}
			index += 4
		case 8:
			leftValue := getAddressOrValue(memory, index+1, modes[len(modes)-1])
			rightValue := getAddressOrValue(memory, index+2, modes[len(modes)-2])
			if leftValue == rightValue {
				setAddressOrValue(memory, storageAddress, 1)
			} else {
				setAddressOrValue(memory, storageAddress, 0)
			}
			index += 4
		default:
			index = len(memory) - 1
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
	// lines := []string{"3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99"}
	linesAsArray := strings.Split(lines[0], ",")

	instructions := instructionsFromArray(linesAsArray)
	runInstructions(instructions)

	// noun, verb := findNounAndVerb(instructions)
	// fmt.Println(noun, verb)
}
