package main

import "fmt"

func readFloatValue(argName string, val *float64) {
	fmt.Print("Enter the value for ", argName, ": ")
	_, err := fmt.Scanf("%f", val)
	if err != nil {
		panic("Unable to read float value for " + argName + " user input")
	}
}

func main() {
	var val float64
	readFloatValue("val", &val)
	fmt.Println("Value is:", val)
}
