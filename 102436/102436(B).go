package main

import (
    "bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	buf := make([]byte, 0, 100100)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, 100100)
	var inputs []string
	//var result int

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		inputs = append(inputs, scanner.Text())
	}

	result := getMostPopularWord(inputs)

	fmt.Println(result)
}

func getMostPopularWord(words []string) int {
	// 1 - letter index, 2 - letter, 3 - repeat count
	ethalonMap := make(map[int]map[uint8]int)

	for _, word := range words {
		for letterIndex := range word {
			letter := word[letterIndex]

			if ethalonMap[letterIndex] != nil {
				ethalonMap[letterIndex][letter]++
			} else {
				m := make(map[uint8]int)
				m[letter]++
				ethalonMap[letterIndex] = m
			}
		}
	}

	var maxRepeatCount int
	bestWord := make(map[int]uint8)
	for letterIndex, letterMap := range ethalonMap {
		maxRepeatCount = 0
		for letter, repeatCount := range letterMap {
			if repeatCount > maxRepeatCount {
				maxRepeatCount = repeatCount
				bestWord[letterIndex] = letter
			}
		}
	}

	diff := 0
	for _, word := range words {
		lenWord := len(word)
		for letterIndex, bestLetter := range bestWord {
			if letterIndex >= lenWord {
				continue
			}

			letter := word[letterIndex]
			if bestLetter != letter {
				diff++
			}
		}
	}

	return diff
}
