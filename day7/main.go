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

func runComputer(inputInstructions []int, inputs []int, startIndex int) (int, int) {
	index := 0

	memory := make([]int, len(inputInstructions))
	copy(memory, inputInstructions)

	inputIndex := 0
	outputs := 0
	for index < len(memory) {
		opCode := memory[index]
		firstAddress := index + 1
		secondAddress := firstAddress + 1
		storageAddress := secondAddress + 1
		modes := []int{0, 0, 0}
		if opCode > 10 && opCode != 99 {
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
		}
		switch opCode {
		case 99:
			index = len(memory)
			return outputs, index
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
			setAddressOrValue(memory, index+1, inputs[inputIndex])
			index += 2
			inputIndex++
		case 4:
			outputs = getAddressOrValue(memory, index+1, modes[2])
			index += 2
			if index > startIndex {
				return outputs, index
			}
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
			index = len(memory)
		}
	}
	fmt.Println("Mem end", memory, index, len(memory))
	return outputs, index
}

func heapPermutation(results [][]int, phases []int, size int) [][]int {
	if size == 1 {
		temp := make([]int, len(phases))
		copy(temp, phases)
		results = append(results, temp)
		return results
	}

	for i := 0; i < size; i++ {
		results = heapPermutation(results, phases, size-1)

		// if size is odd, swap first and last
		// element
		temp := make([]int, len(phases))
		copy(temp, phases)
		if size%2 == 1 {
			phases[0] = phases[size-1]
			phases[size-1] = temp[0]

			// If size is even, swap ith and last
			// element
		} else {
			phases[i] = phases[size-1]
			phases[size-1] = temp[i]
		}
	}
	return results
}

func main() {
	lines := readFileToArray("input.txt")
	linesAsArray := strings.Split(lines[0], ",")

	instructions := instructionsFromArray(linesAsArray)
	maxAmp := -99999
	phases := []int{5, 6, 7, 8, 9}
	results := [][]int{}
	results = heapPermutation(results, phases, 5)
	loop := true

	for p := 0; p < len(results); p++ {
		loop = true
		input1, input2, input3, input4, input5 := []int{results[p][0]}, []int{results[p][1]}, []int{results[p][2]}, []int{results[p][3]}, []int{results[p][4]}
		out1, out2, out3, out4, out5 := 0, 0, 0, 0, 0
		ind1, ind2, ind3, ind4, ind5 := 0, 0, 0, 0, 0
		for loop {
			input1 = append(input1, out5)
			out1, ind1 = runComputer(instructions, input1, ind1)
			input2 = append(input2, out1)
			out2, ind2 = runComputer(instructions, input2, ind2)
			input3 = append(input3, out2)
			out3, ind3 = runComputer(instructions, input3, ind3)
			input4 = append(input4, out3)
			out4, ind4 = runComputer(instructions, input4, ind4)
			input5 = append(input5, out4)
			out5, ind5 = runComputer(instructions, input5, ind5)
			if ind5 < len(instructions) {
				if out5 > maxAmp {
					maxAmp = out5
				}
			} else {
				loop = false
			}

		}

	}
	fmt.Println(maxAmp)

}
