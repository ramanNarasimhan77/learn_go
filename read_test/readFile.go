package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("/Users/ramann/Documents/experiments/go/read/names.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
