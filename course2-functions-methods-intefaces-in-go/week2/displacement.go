package main

import (
	"fmt"
	"strings"
)

// GenDisplaceFn calculates displacement s as a function of time t,
// acceleration a, initial velocity vo, and initial displacement  using formula
// s =½ a t2 + vot + so
func GenDisplaceFn(a float64, v0 float64, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return .5*a*t*t + v0*t + s0
	}
}

func readFloatValue(argName string) float64 {
	var argValue float64
	fmt.Print("Enter the value for ", argName, ": ")
	_, err := fmt.Scanf("%f", &argValue)
	if err != nil {
		panic("Unable to read float value for " + argName + " user input")
	}
	return argValue
}

func main() {
	acceleration := readFloatValue("acceleration (a)")
	initialVelocity := readFloatValue("initial velocity (v0)")
	initialDisplacement := readFloatValue("initial displacement (s0)")
	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)
	rerun := true
	var continueLoop string
	for rerun {
		time := readFloatValue("time (t)")
		fmt.Println("Displacement after ", time, " seconds is:", fn(time))
		fmt.Print("Do you want to calculate displacement for another value of time? Enter yes to continue :")
		fmt.Scanf("%s", &continueLoop)
		rerun = strings.EqualFold(continueLoop, "yes")
		if !rerun {
			fmt.Println("Goodbye..")
		}
	}
}
