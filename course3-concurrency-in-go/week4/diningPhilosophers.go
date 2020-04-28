package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ordinalMap = map[int]string{
	1: "first",
	2: "second",
	3: "third",
	4: "fourth",
}

// Chopstick is a struct with a mutex
type Chopstick struct{ sync.Mutex }

// Philosopher has an id and 2 chopsticks
type Philosopher struct {
	id             int
	leftChopstick  *Chopstick
	rightChopstick *Chopstick
}

// GetChopsticks - Philosopher gets permission from host by reading a value from channel c
// and then tries to get the chopsticks
// channel c can hold only upto 2 values. Philosophers get blocked if channel has no data
func (p Philosopher) GetChopsticks(c *chan int) {
	<-*c
	//fmt.Println("Philosopher", p.id, "has got approval from host to start eating")
	p.leftChopstick.Lock()
	p.rightChopstick.Lock()
}

// ReleaseChopsticks - Philosopher releases the chopsticks
func (p Philosopher) ReleaseChopsticks() {
	//fmt.Println("Philosopher", p.id, "has finished eating")
	fmt.Printf("\nfinishing eating %d", p.id)
	p.leftChopstick.Unlock()
	p.rightChopstick.Unlock()
}

// Eat - philosopher acquires the chopsticks and eats for upto 25 millisec
func (p Philosopher) Eat(c *chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 3; i++ {
		p.GetChopsticks(c)
		//fmt.Println("Philosopher", p.id, "is eating for the", ordinalMap[i], "time")
		fmt.Printf("\nstarting to eat %d", p.id)
		time.Sleep(time.Duration(rand.Intn(25)) * time.Millisecond)
		p.ReleaseChopsticks()
	}
}

//Host - writes two values to the channel c, philosophers should read from c and get a value to start eating
func Host(c *chan int) {
	for {
		if len(*c) == 0 {
			*c <- 1
			*c <- 1
			time.Sleep(25 * time.Millisecond)
		}
	}
}

func main() {
	const noOfPhilosophers = 5
	var wg sync.WaitGroup
	channel := make(chan int, 2)
	chopsticks := make([]*Chopstick, noOfPhilosophers)
	philosophers := make([]*Philosopher, noOfPhilosophers)

	for i := 0; i < len(chopsticks); i++ {
		chopsticks[i] = new(Chopstick)
	}

	go Host(&channel)

	for i := 0; i < len(philosophers); i++ {
		philosophers[i] = &Philosopher{i + 1, chopsticks[i], chopsticks[(i+1)%5]}
		wg.Add(1)
		go philosophers[i].Eat(&channel, &wg)
	}
	wg.Wait()
}
