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
	n, _ := strconv.Atoi(scanner.Text())

	var data []string
	var l1, r1, l2, r2, d1, d2 int
	for testNum := 0; testNum < n; testNum++ {
		scanner.Scan()
		data = strings.Split(scanner.Text(), " ")
		l1, _ = strconv.Atoi(data[0])
		r1, _ = strconv.Atoi(data[1])
		l2, _ = strconv.Atoi(data[2])
		r2, _ = strconv.Atoi(data[3])

		scanner.Scan()
		data = strings.Split(scanner.Text(), " ")
		s1, _ := strconv.Atoi(data[0])
		d1, _ = strconv.Atoi(data[1])
		s2, _ := strconv.Atoi(data[2])
		d2, _ = strconv.Atoi(data[3])

		if s1+d1 <= s2 {
			fmt.Println(s1, s2)
			continue
		}

		if s1 >= s2+d2 {
			fmt.Println(s1, s2)
			continue
		}

		rList := make([]int, 2)
		rList[0], rList[1] = -1, -1

		var sDiv, minDiv int
		minDiv = abs(r1-l1) + abs(r2-l2)

		for i := l1; i+d1 <= r1; i++ {
			for j := l2; j+d2 <= r2; j++ {
				if i+d1 <= j || i >= j+d2 {
					sDiv = abs(s1-i) + abs(s2-j)

					if minDiv > sDiv {
						rList[0], rList[1] = i, j
						minDiv = sDiv
					}
				}
			}

			if minDiv == 0 {
				break
			}
		}

		fmt.Println(rList[0], rList[1])
	}
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}

	return n
}
