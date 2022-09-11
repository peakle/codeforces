package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var st []int64

func init() {
	for i := 9; i >= 0; i-- {
		st = append(st, int64(math.Pow(10, float64(i))))
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	val, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < val; i++ {
		scanner.Scan()
		m, _ := strconv.ParseInt(scanner.Text(), 10, 64)

		for _, k := range st {
			if d := m - k; d >= 0 && d < m {
				fmt.Println(d)
				break
			}
		}
	}
}
