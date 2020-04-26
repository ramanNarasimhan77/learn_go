package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal Type defining features.
type Animal struct {
	food       string
	locomotion string
	noise      string
}

// Init initializes an Animal
func (animal *Animal) Init(food, locomotion, noise string) {
	animal.food = food
	animal.locomotion = locomotion
	animal.noise = noise
}

// Eat returns animal.food
func (animal *Animal) Eat() string {
	return animal.food
}

// Move returns animal.locomotion
func (animal *Animal) Move() string {
	return animal.locomotion
}

// Speak returns animal.noise
func (animal *Animal) Speak() string {
	return animal.noise
}

// GetInfo invokes Eat. Sleep or Move based on query parameter
func (animal *Animal) GetInfo(query string) string {
	switch query {
	case "eat":
		return animal.Eat()
	case "speak":
		return animal.Speak()
	case "move":
		return animal.Move()
	default:
		return "Animal can only eat, speak or move. Query " + query + " is undefined"
	}
}

func readUserInput(input *string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	*input = scanner.Text()
}

func main() {
	var cow, bird, snake Animal
	cow.Init("grass", "walk", "moo")
	bird.Init("worms", "fly", "peep")
	snake.Init("mice", "slither", "hiss")
	var animals = map[string]Animal{
		"cow":   cow,
		"bird":  bird,
		"snake": snake,
	}

	continueLoop := true
	fmt.Println(" Starting Program..Enter exit to quit")
	var userQuery, animalName, query string

	for continueLoop {
		fmt.Print("> ")
		readUserInput(&userQuery)
		if strings.EqualFold("exit", userQuery) {
			fmt.Println("GoodBye")
			continueLoop = false
		} else {
			queryParts := strings.Split(userQuery, " ")
			if len(queryParts) != 2 {
				fmt.Println("Invalid query. Query should have two fields. Format:  cow|bird|snake eat|move|speak")
				continue
			}
			animalName = queryParts[0]
			query = queryParts[1]

			animal, keyPresent := animals[animalName]
			if keyPresent {
				fmt.Println(animal.GetInfo(query))
			} else {
				fmt.Println("Animal", animalName, "is invalid. Valid animals are cow, bird, snake")
			}
		}
	}
}
