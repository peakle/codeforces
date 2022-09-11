package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var maxSize int

func init() {
	maxSize = 2 * int(math.Pow(10, 5)) + 3
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, maxSize), maxSize)
	scanner.Scan()

	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		t := scanner.Text()

		c := 0
		uniq := make(map[rune]struct{}, 3)

		for _, r := range t {
			uniq[r] = struct{}{}

			if len(uniq) > 3 {
				c++
				uniq = map[rune]struct{}{r: {}}
			}
		}

		if len(uniq) > 0 {
			c++
		}

		fmt.Println(c)
	}
}
