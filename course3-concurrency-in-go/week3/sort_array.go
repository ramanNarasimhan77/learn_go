package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func readArray(sli *[]int) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the numbers on same line separated by space (> 4 numbers)")
	fmt.Print("> ")
	scanner.Scan()
	userInputStr := strings.TrimRight(scanner.Text(), " ")
	inputValues := strings.Split(userInputStr, " ")
	for _, inputValue := range inputValues {
		inputInt, err := strconv.Atoi(inputValue)
		if err != nil {
			fmt.Println("Could not convert userinput", inputValue, " to a valid int. Error is", err.Error())
		}
		*sli = append(*sli, inputInt)
	}
}

func sortPartition(sli *[]int, grpID int, wg *sync.WaitGroup) {
	defer (*wg).Done()
	msgPrefix := "[grp" + fmt.Sprintf("%d", grpID) + "]"
	fmt.Println(msgPrefix, "Sorting", *sli)
	sort.Ints(*sli)
	fmt.Println(msgPrefix, "Sorted result", *sli)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	const noOfPartitions = 4
	var waitGroup sync.WaitGroup
	var arr = make([]int, 0, noOfPartitions)
	readArray(&arr)
	var arrlen = len(arr)
	var noOfEleInAGrp = max(1, arrlen/noOfPartitions)
	var lowerBound, upperBound int
	fmt.Println("Starting Execution Now")
	for i := 0; i < noOfPartitions; i++ {
		lowerBound = min(i*noOfEleInAGrp, arrlen)
		upperBound = min((i+1)*noOfEleInAGrp, arrlen)
		//fmt.Println("lowerbound:", lowerBound, "upperBound:", upperBound)
		if i == noOfPartitions-1 && upperBound < arrlen { // in case no of elements in arr is not exactly divisible by noOfPartitions, add spillOver elements to lastSubArray
			upperBound = arrlen
		}
		//fmt.Println("lowerbound:", lowerBound, "upperBound:", upperBound)
		sli := arr[lowerBound:upperBound]
		waitGroup.Add(1)
		go sortPartition(&sli, i, &waitGroup)
	}
	waitGroup.Wait()
	fmt.Println("Array with sorted partitions", arr)
	sort.Ints(arr)
	fmt.Println("Sorted array", arr)
}
