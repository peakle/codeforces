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
	buf := make([]byte, 1000000000)
	scanner.Buffer(buf, len(buf))

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	var data []string
	var l1, r1, l2, r2, s1, d1, s2, d2, lm, rm, ld int
	for i := 0; i < n; i++ {
		scanner.Scan()
		data = strings.Split(scanner.Text(), " ")
		l1, _ = strconv.Atoi(data[0])
		r1, _ = strconv.Atoi(data[1])
		l2, _ = strconv.Atoi(data[2])
		r2, _ = strconv.Atoi(data[3])

		scanner.Scan()
		data = strings.Split(scanner.Text(), " ")
		s1, _ = strconv.Atoi(data[0])
		d1, _ = strconv.Atoi(data[1])
		s2, _ = strconv.Atoi(data[2])
		d2, _ = strconv.Atoi(data[3])

		lm = getMin(l1, l2)
		rm = getMax(r1, r2)
		ld = lm - rm

		if ld < (d1 + d2) {
			fmt.Println(-1, -1)
			continue
		}

		fmt.Println()
	}
}

func getMax(f, t int) int {
	if f > t {
		return f
	}

	return t
}

func getMin(f, t int) int {
	if f > t {
		return t
	}

	return f
}
