package main

import (
	"fmt"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func main() {
	input := `357253-892942`
	inputs := strings.Split(input, "-")
	start, serr := strconv.Atoi(inputs[0])
	finish, ferr := strconv.Atoi(inputs[1])
	meetCriteria := 0
	check(serr)
	check(ferr)
	/**
	Loop over input range
	**/
	for x := start; x <= finish; x++ {
		onlyIncreases := true
		hasDouble := false
		doubleDigit := -1
		// Loop over each digit in our password
		// Need lookahead
		xAsArray := strconv.Itoa(x)
		for i := 0; i < len(xAsArray); i++ {
			digit, err := strconv.Atoi(string(xAsArray[i]))
			check(err)

			if i+1 < len(xAsArray) {
				oneAhead, errL := strconv.Atoi(string(xAsArray[i+1]))
				check(errL)
				if i+2 < len(xAsArray) {
					twoAhead, errL2 := strconv.Atoi(string(xAsArray[i+2]))
					check(errL2)

					if digit == twoAhead && digit == oneAhead {
						doubleDigit = digit
					}
					if digit != doubleDigit && digit == oneAhead && digit != twoAhead {
						hasDouble = true
					}
				} else {
					if digit != doubleDigit && digit == oneAhead {
						hasDouble = true
					}
				}

				if digit > oneAhead {
					onlyIncreases = false
				}
			}
		}
		if onlyIncreases && hasDouble {
			meetCriteria++
		}
	}

	fmt.Println(meetCriteria)
}
