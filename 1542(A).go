package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// четных и нечетных должно быть одинаково
// повторяющихся должно быть не больше количества пар
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	testCount, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < testCount; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		input := scanner.Text()

		inputList := make([]int, 0, 100)
		for _, p := range strings.Split(input, " ") {
			o, _ := strconv.Atoi(p)
			inputList = append(inputList, o)
		}

		var sum, even, odd int
		var repeatedCount = make(map[int]int, len(inputList))
		for _, s := range inputList {
			sum += s
			if s%2 == 0 {
				odd++
			} else {
				even++
			}

			repeatedCount[s]++
		}
		var mostRepeated int
		for _, c := range repeatedCount {
			if c > mostRepeated {
				mostRepeated = c
			}
		}

		if mostRepeated > n || odd != even {
			fmt.Println("No")
			continue
		}

		if (n%2 == 0 && sum%2 != 0) || (n%2 != 0 && sum%2 == 0) {
			fmt.Println("No")
			continue
		}

		fmt.Println("Yes")
	}
}
