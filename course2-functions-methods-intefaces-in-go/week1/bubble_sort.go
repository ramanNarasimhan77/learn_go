package main

import (
	"errors"
	"fmt"
	"strconv"
)

func swap(numbers []int, index int) {
	temp := numbers[index]
	numbers[index] = numbers[index+1]
	numbers[index+1] = temp
}

// BubbleSort takes a slice of numbers and sorts it.
func BubbleSort(numbers []int) {
	cnt := len(numbers)
	fmt.Println("Numbers before sorting", numbers)
	for i := 0; i < cnt; i++ {
		for j := 0; j < cnt-i-1; j++ {
			if numbers[j+1] < numbers[j] {
				swap(numbers, j)
			}
		}
	}
}

func readNumbers() []int {
	var i int
	var numbers = make([]int, 0)
	for i < 10 {
		fmt.Println("Enter a number (to exit enter exit):")
		var input string
		_, err := fmt.Scanf("%s", &input)
		if err != nil {
			panic(errors.New("Failed to read number"))
		}
		if input == "exit" {
			break
		} else {
			inputInt, err := strconv.Atoi(input)
			if err != nil {
				panic(errors.New("Failed to convert " + input + " to number"))
			}
			numbers = append(numbers, inputInt)
		}
		i++
	}
	if i == 10 {
		fmt.Println("Program has Read 10 numbers.")
	}
	return numbers
}

func main() {
	inputNums := readNumbers()
	BubbleSort(inputNums)
	fmt.Println("Numbers after sorting", inputNums)
}
