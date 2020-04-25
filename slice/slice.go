package main

import (
	"fmt"
	"sort"
	"strconv"
)

func readUserInput() string {
	var input string
	fmt.Printf("\n\nEnter a number (to exit enter X):")
	fmt.Scanf("%s", &input)
	return input
}

func addToSlice(sli []int, intValue int) []int {
	sli = append(sli, intValue)
	sort.Ints(sli)
	fmt.Printf("\nAdded element %d", intValue)
	return sli
}

func main() {
	initialCapacity := 3
	sli := make([]int, 0, initialCapacity) // initializes an empty slice with capacity=3

	for {
		inputStr := readUserInput()
		if inputStr == "X" {
			fmt.Println("Goodbye!!")
			break
		} else {
			inputInt, err := strconv.Atoi(inputStr)
			if err != nil {
				fmt.Printf("\n%s is not a valid input.\nTry entering an integer or enter X to exit", inputStr)
			} else {
				sli = addToSlice(sli, inputInt)
				fmt.Printf("\nSorted slice is: %v", sli)
			}
		}
	}
}
