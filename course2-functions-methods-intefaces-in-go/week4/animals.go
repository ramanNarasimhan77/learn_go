package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal interface defines methods applicable to animals.
type Animal interface {
	Eat()
	Move()
	Speak()
}

/*
Struct Definitions
*/

//Cow type
type Cow struct{}

//Snake Type
type Snake struct{}

//Bird Type
type Bird struct{}

/*
Functions for Type Cow
*/

// Eat for Type Cow
func (cow *Cow) Eat() {
	fmt.Println("grass")
}

//Move for Type Cow
func (cow *Cow) Move() {
	fmt.Println("walk")
}

// Speak for type Cow
func (cow *Cow) Speak() {
	fmt.Println("moo")
}

/*
Functions for type Bird
*/

// Eat for Type Bird
func (bird *Bird) Eat() {
	fmt.Println("worms")
}

// Move for type Bird
func (bird *Bird) Move() {
	fmt.Println("fly")
}

// Speak for type Bird
func (bird *Bird) Speak() {
	fmt.Println("peep")
}

/*
Functions for Type Snake
*/

// Eat for Type Snake
func (snake *Snake) Eat() {
	fmt.Println("mice")
}

// Move for type snake
func (snake *Snake) Move() {
	fmt.Println("slither")
}

// Speak for type snake
func (snake *Snake) Speak() {
	fmt.Println("hsss")
}

func createNewAnimal(animalType string, animalName string, animals *map[string]Animal) {
	created := true
	switch animalType {
	case "cow":
		(*animals)[animalName] = new(Cow)
	case "snake":
		(*animals)[animalName] = new(Snake)
	case "bird":
		(*animals)[animalName] = new(Bird)
	default:
		created = false
		fmt.Println("Cannot create an animal of type", animalType)
	}
	if created {
		fmt.Println("Created!")
	}
}

func queryAnimalInfo(query string, animalName string, animals *map[string]Animal) {
	animal, ok := (*animals)[animalName]
	if !ok {
		fmt.Println("Animal named", animalName, "is not known. Please create it first using newanimal command")
		return
	}
	switch query {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("Action", query, "is not defined. Valid actions are eat|move|speak")
	}
}

func main() {
	fmt.Println("Starting the program. To quit enter exit")
	continueLoop := true
	var command, animalName, cmdInfo string
	animals := make(map[string]Animal)
	var scanner = bufio.NewScanner(os.Stdin)
	for continueLoop {
		fmt.Print("> ")
		scanner.Scan()
		userInput := scanner.Text()
		if strings.EqualFold(userInput, "exit") {
			fmt.Println("Goodbye!!")
			continueLoop = false
		} else {
			inputFields := strings.Split(userInput, " ")
			if len(inputFields) != 3 {
				fmt.Println("User input should have 3 parts newanimal|query animalName animalType|queryType")
				continue
			}
			command, animalName, cmdInfo = inputFields[0], inputFields[1], inputFields[2]
			switch command {
			case "newanimal":
				createNewAnimal(cmdInfo, animalName, &animals)
			case "query":
				queryAnimalInfo(cmdInfo, animalName, &animals)
			}
		}
	}
}
