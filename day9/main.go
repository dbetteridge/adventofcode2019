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

func instructionsFromArray(array []string) []int64 {
	instructions := []int64{}
	for _, x := range array {
		singleInt64, err := strconv.ParseInt(x, 10, 64)
		check(err)
		singleInt := int64(singleInt64)
		instructions = append(instructions, singleInt)
	}
	return instructions
}

func getAddressOrValue(memory []int64, index int64, mode int, base int64) int64 {
	if mode == 0 {
		return memory[memory[index]]
	}
	if mode == 2 {
		return memory[base+memory[index]]
	}
	return memory[index]
}

func setAddressOrValue(memory []int64, index int64, value int64, mode int, base int64) {
	if mode == 2 {
		memory[base+memory[index]] = value
	} else {
		memory[memory[index]] = value
	}
}

func runComputer(inputInstructions []int64, inputs []int64, startIndex int64) (int64, int64) {
	var index int64 = 0

	memory := make([]int64, len(inputInstructions)+999999)
	copy(memory, inputInstructions)

	var inputIndex int64 = 0
	var outputs int64 = 0
	var base int64 = 0
	for index < int64(len(memory)) {
		opCode := memory[index]
		firstAddress := index + 1
		secondAddress := firstAddress + 1
		storageAddress := secondAddress + 1
		modes := []int{0, 0, 0}
		if opCode > 10 && opCode != 99 {
			imOpCode := strconv.Itoa(int(opCode))
			opCode = 0
			modeCount := 5 - len(imOpCode)
			for i, c := range imOpCode {
				intC, err := strconv.Atoi(string(c))
				check(err)
				if i == len(imOpCode)-1 || i == len(imOpCode)-2 {
					opCode += int64(intC)
				} else {
					modes[modeCount] = intC
					modeCount++
				}
			}
		}
		switch opCode {
		case 99:
			index = int64(len(memory))
			return outputs, index
		case 1:
			leftValue := getAddressOrValue(memory, firstAddress, modes[len(modes)-1], base)
			rightValue := getAddressOrValue(memory, secondAddress, modes[len(modes)-2], base)
			setAddressOrValue(memory, storageAddress, leftValue+rightValue, modes[len(modes)-3], base)
			index += 4
		case 2:
			leftValue := getAddressOrValue(memory, firstAddress, modes[len(modes)-1], base)
			rightValue := getAddressOrValue(memory, secondAddress, modes[len(modes)-2], base)
			setAddressOrValue(memory, storageAddress, leftValue*rightValue, modes[len(modes)-3], base)
			index += 4
		case 3:
			setAddressOrValue(memory, index+1, inputs[inputIndex], modes[len(modes)-1], base)
			index += 2
			inputIndex++
		case 4:
			outputs = getAddressOrValue(memory, index+1, modes[2], base)
			index += 2
			if index > startIndex {
				fmt.Println(outputs)
			}
		case 5:
			leftValue := getAddressOrValue(memory, index+1, modes[len(modes)-1], base)
			rightValue := getAddressOrValue(memory, index+2, modes[len(modes)-2], base)
			if leftValue != 0 {
				index = rightValue
			} else {
				index += 3
			}
		case 6:
			leftValue := getAddressOrValue(memory, index+1, modes[len(modes)-1], base)
			rightValue := getAddressOrValue(memory, index+2, modes[len(modes)-2], base)
			if leftValue == 0 {
				index = rightValue
			} else {
				index += 3
			}
		case 7:
			leftValue := getAddressOrValue(memory, index+1, modes[len(modes)-1], base)
			rightValue := getAddressOrValue(memory, index+2, modes[len(modes)-2], base)
			if leftValue < rightValue {
				setAddressOrValue(memory, storageAddress, 1, modes[len(modes)-3], base)
			} else {
				setAddressOrValue(memory, storageAddress, 0, modes[len(modes)-3], base)
			}
			index += 4
		case 8:
			leftValue := getAddressOrValue(memory, index+1, modes[len(modes)-1], base)
			rightValue := getAddressOrValue(memory, index+2, modes[len(modes)-2], base)
			if leftValue == rightValue {
				setAddressOrValue(memory, storageAddress, 1, modes[len(modes)-3], base)
			} else {
				setAddressOrValue(memory, storageAddress, 0, modes[len(modes)-3], base)
			}
			index += 4
		case 9:
			leftValue := getAddressOrValue(memory, firstAddress, modes[len(modes)-1], base)
			base += leftValue
			index += 2
		default:
			index = int64(len(memory))
		}
	}
	fmt.Println("Mem end", memory, index, len(memory))
	return outputs, index
}

func main() {
	lines := readFileToArray("input.txt")
	// lines := []string{"109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"}
	linesAsArray := strings.Split(lines[0], ",")

	instructions := instructionsFromArray(linesAsArray)
	runComputer(instructions, []int64{2}, 0)

}
