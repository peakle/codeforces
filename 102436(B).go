package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var inputs []string
	var result int

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		inputs = append(inputs, scanner.Text())
	}

	result = getMinDiffWord(inputs)

	fmt.Println(result)
}

func getMinDiffWord(words []string) int {
	minDiff := 100000

	for _, wordA := range words {
		diff := 0
		lenWordA := len(wordA)

		for _, wordB := range words {
			if wordA != wordB {
				lenWordB := len(wordB)

				for indexWordA := range wordA {
					if indexWordA >= lenWordB {
						continue
					}

					letterA := wordA[indexWordA]
					letterB := wordB[indexWordA]

					if letterA != letterB {
						diff++
					}
				}

				if lenWordA <= lenWordB {
					diff += lenWordB - lenWordA
				}
			}
		}

		if diff < minDiff && diff != 0 {
			minDiff = diff
		}
	}

	return minDiff
}
