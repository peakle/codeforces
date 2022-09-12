package main

import (
    "bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	input := scanner.Text()

	testCount, _ := strconv.Atoi(input)
	var inputList []string
	var candyCount, childCount int

	for i := 0; i < testCount; i++ {
		scanner.Scan()
		input = scanner.Text()

		inputList = strings.Split(input, " ")
		candyCount, _ = strconv.Atoi(inputList[0])
		childCount, _ = strconv.Atoi(inputList[1])

		candyDiff := candyCount % childCount;
		if candyDiff == 0 {
			fmt.Println(candyCount)
			continue
		}

		baseCandyCount := candyCount - candyDiff
		prevChildCount := childCount / 2

		// нужно чтоб ответ не был больше количества конфет ы
		fmt.Println(min(candyCount, baseCandyCount+prevChildCount))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
