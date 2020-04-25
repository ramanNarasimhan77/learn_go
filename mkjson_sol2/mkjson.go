package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func readUserInput1() string {
	var input = make([]string, 0, 1)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}
	return strings.Join(input, "\n")
}

func readUserInput(isMultiline bool) string {
	scanner := bufio.NewScanner(os.Stdin)
	var userInput string
	if isMultiline {
		var userInputs = make([]string, 0, 1)
		stillReading := true
		for stillReading {
			scanner.Scan()
			line := scanner.Text()
			if strings.EqualFold(line, "Done") {
				stillReading = false
				userInput = strings.Join(userInputs, "\n")
				break
			} else {
				userInputs = append(userInputs, line)
			}
		}
	} else {
		scanner.Scan()
		userInput = scanner.Text()
	}
	return userInput
}

func main() {
	addressMap := make(map[string]string)

	fmt.Println("Enter your name:")
	addressMap["name"] = readUserInput(false)
	fmt.Println("Enter your address")
	addressMap["address"] = readUserInput(false)

	jsonData, _ := json.Marshal(addressMap)

	fmt.Println("Json Data is" + string(jsonData) + "\n")
}
