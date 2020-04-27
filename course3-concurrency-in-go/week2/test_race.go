package main

import (
	"fmt"
	"sync"
	"time"
)

// Counter is a counter
var Counter int = 0

//Wait is a WaitGroup
var Wait sync.WaitGroup

/*
IncCounter  assigns the counter to a temp variable,
sleeps for 1 Nanosecond
increments the temp variable
reassigns temp variable back to Counter

Counter is a global variable which is Read & Set in this function

*/
func IncCounter(id int) {
	for i := 0; i < 2; i++ {
		val := Counter
		time.Sleep(1 * time.Nanosecond)
		val++
		Counter = val
		fmt.Println("routine", id, "in loop", i+1, "set counter value to", Counter, "at", time.Now().Format(time.StampNano))
	}
	Wait.Done()
}

func main() {

	/*

		In case the IncCounter function is invoked twice sequentially, the final value for counter would be 4

		the for loop below uses GoRoutines to spawns 2 threads that invoke the IncCounter function
		the execution of the goroutines would be interleaved and
		and each goroutine execution is non-deterministic

		Also the func IncCounter READS and SETS a shared variable Counter
		Depending on when the GoRoutine reads the Counter, it would get a different value for the counter

		Consider the scenario where
		GoRoutine1 reads value of Counter (0) into temp variable val and Sleeps for 1 Nanosecond
		Before it wakes up, GoRoutine2 also reads Counter value (0) and sets temp variable val and sleeps for 1 Nanosecond
		Now GoRoutine1 wakes up, increments val to 1 and assigns to Counter. Now Counter=1
		After this GoRoutine2 wakes up, it has val still set to 0 anf so it increments val to 1 and assigns to Counter. Now Counter is again=1
		The update made by GoRoutine1 is lost

		Another Scenario is
		GoRoutine1 completes loop iteration i=0, sets Counter to 1
		GoRoutine2 starts loop iteration i=0 after this, and so it sets Counter to 2
		Before GoRoutine2 can increment Counter to 2 say GoRoutine1 starts iteration i=1, then it would start with Counter value of 1 and increment it to 2
		GoRoutine2 iteration i=1 starts after GoRoutine1 iteration i=1 and it increments Counter from 2 to 3

		Go is equipped with an inbuilt race detector to detect these race conditions
		Execute below command to see the code snippet causing race condition
		You can see race is detected at line 26 and 29 where Counter is READ and WRITTEN to respectively

		go run -race test_race.go

		==================
		WARNING: DATA RACE
		Write at 0x000001283580 by goroutine 8:
		main.IncCounter()
			/Users/ramann/Documents/experiments/go/course3-concurrency-in-go/week2/test_race.go:29 +0x87

		Previous read at 0x000001283580 by goroutine 7:
		main.IncCounter()
			/Users/ramann/Documents/experiments/go/course3-concurrency-in-go/week2/test_race.go:26 +0x5e

		Goroutine 8 (running) created at:
		main.main()
			/Users/ramann/Documents/experiments/go/course3-concurrency-in-go/week2/test_race.go:115 +0x75

		Goroutine 7 (running) created at:
		main.main()
			/Users/ramann/Documents/experiments/go/course3-concurrency-in-go/week2/test_race.go:115 +0x75
		==================
		==================
		WARNING: DATA RACE
		Write at 0x000001283580 by goroutine 7:
		main.IncCounter()
			/Users/ramann/Documents/experiments/go/course3-concurrency-in-go/week2/test_race.go:29 +0x87

		Previous read at 0x000001283580 by goroutine 8:
		main.IncCounter()
			/Users/ramann/Documents/experiments/go/course3-concurrency-in-go/week2/test_race.go:26 +0x5e

		Goroutine 7 (running) created at:
		main.main()
			/Users/ramann/Documents/experiments/go/course3-concurrency-in-go/week2/test_race.go:115 +0x75

		Goroutine 8 (running) created at:
		main.main()
			/Users/ramann/Documents/experiments/go/course3-concurrency-in-go/week2/test_race.go:115 +0x75
		==================
		routine 1 in loop 1 set counter value to 1 at Apr 27 20:29:12.620700000
		routine 2 in loop 1 set counter value to 1 at Apr 27 20:29:12.620645000
		routine 2 in loop 2 set counter value to 2 at Apr 27 20:29:12.621093000
		routine 1 in loop 2 set counter value to 2 at Apr 27 20:29:12.621190000
		Final Counter: 2
		Found 2 data race(s)
		exit status 66
	*/

	for routine := 1; routine <= 2; routine++ {

		Wait.Add(1)
		go IncCounter(routine)
	}

	Wait.Wait()
	fmt.Printf("Final Counter: %d\n", Counter)

}
