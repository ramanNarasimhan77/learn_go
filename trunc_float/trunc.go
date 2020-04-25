package main

import (
	"fmt"
)

func main() {
	var floatNum float64
	var intNum int64
	fmt.Printf("Enter a floating point number:")
	fmt.Scanf("%f", &floatNum)
	fmt.Printf("You have entered:%f\n", floatNum)
	intNum = int64(floatNum)
	fmt.Printf("Integer value for input is %d\n", intNum)
}
