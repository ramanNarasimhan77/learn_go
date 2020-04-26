package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string:")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while reading input string:" + err.Error())
	}
	input = strings.TrimRight(input, "\r\n")
	inputLowerCase := strings.ToLower(input)
	if strings.HasPrefix(inputLowerCase, "i") &&
		strings.HasSuffix(inputLowerCase, "n") &&
		strings.Contains(inputLowerCase, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
