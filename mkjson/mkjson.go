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
	reader := bufio.NewReader(os.Stdin)
	var userInput string
	if isMultiline {
		var userInputs = make([]string, 0, 1)
		stillReading := true
		for stillReading {
			line, _ := reader.ReadString('\n')
			line = strings.TrimSuffix(line, "\n")
			if strings.EqualFold(line, "~") {
				stillReading = false
				userInput = strings.Join(userInputs, "\n")
				break
			} else {
				userInputs = append(userInputs, line)
			}
		}
	} else {
		userInput, _ = reader.ReadString('\n')
		userInput = strings.TrimSuffix(userInput, "\n")
	}
	return userInput
}

func main() {
	addressMap := make(map[string]string)

	fmt.Println("Enter your name:")
	addressMap["name"] = readUserInput(false)
	fmt.Println("Enter your address line by line. Once done, Type ~ in a newline and press enter")
	addressMap["address"] = readUserInput(true)

	jsonData, _ := json.Marshal(addressMap)

	fmt.Println("Json Data is" + string(jsonData) + "\n")
}
