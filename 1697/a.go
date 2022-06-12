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

	c, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < c; i++ {
		scanner.Scan()
		sub := strings.Split(scanner.Text(), " ")
		//_, _ := strconv.Atoi(sub[0])
		m, _ := strconv.ParseInt(sub[1], 10, 64)

		var dists []int
		scanner.Scan()
		for _, v := range strings.Split(scanner.Text(), " ") {
			d, _ := strconv.Atoi(v)
			dists = append(dists, d)
		}

		var sum int64
		for _, v := range dists {
			sum += int64(v)
		}

		r := sum - m
		if r > 0 {
			fmt.Println(r)
		} else {
			fmt.Println(0)
		}
	}
}
