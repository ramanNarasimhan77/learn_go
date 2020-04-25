package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// PersonName struct
type PersonName struct {
	fName string
	lName string
}

func trimLength(s string, n int) string {
	rs := []rune(s)
	return string(rs[:20])
}

func processLine(line string) (string, string) {
	fields := strings.Split(line, " ")
	if len(fields) < 2 {
		err := fmt.Errorf("Record %s does not have two columns", line)
		check(err)
	}
	return trimLength(fields[0], 20), trimLength(fields[1], 20)
}

func check(e error) {
	if e != nil {
		fmt.Println("Error is:", e.Error())
		panic(e)
	}
}

func fileExists(filename string) bool {
	fmt.Println("Searching for file:", filename, " length:", len(filename))
	info, err := os.Stat(filename)
	if err != nil {
		fmt.Errorf("Failed to find file", err.Error())
		return false
	}
	return !info.IsDir()
}

func trimSuffix(str string) string {
	var suffix string
	if strings.HasSuffix(str, "\r\n") {
		suffix = "\r\n"
	} else if strings.HasSuffix(str, "\n") {
		suffix = "\n"
	}
	return strings.TrimRight(str, suffix)
}

func readPersonNamesFromFile(pathToFile string) []PersonName {
	var personNames = make([]PersonName, 0)
	file, fileErr := os.Open(pathToFile)
	if fileErr != nil {
		println(fileErr, "unable to open %s", pathToFile)
	}
	check(fileErr)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		firstName, lastName := processLine(line)
		personNames = append(personNames, PersonName{firstName, lastName})
	}
	return personNames
}

func main() {
	fmt.Println("Enter the file path:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	filePath := scanner.Text()
	if fileExists(filePath) {
		println("Found file.. ", filePath)
		personNames := readPersonNamesFromFile(filePath)
		for i, personName := range personNames {
			fmt.Println("Entry ", i+1, "\t FirstName:", personName.fName, " LastName:", personName.lName)
		}
	} else {
		err := fmt.Errorf("Fatal Error... File %s does not exist", filePath)
		check(err)
	}

}
