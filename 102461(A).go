package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	var result string

	strLen := len(input) - 1
	for index, letter := range input {
		if isSymbol(string(letter)) {
			if strLen > index {
				result += " " + input[index:index+1] + " "
			} else {
				result += " " + input[index:strLen+1]
			}
		} else {
			result += input[index : index+1]
		}
	}

	fmt.Println(result)
}

func isSymbol(l string) bool {
	switch l {
	case "+", "-", "*", "/":
		return true
	}

	return false
}
