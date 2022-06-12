package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const maxTokenSize = 262144 * 10

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, maxTokenSize), maxTokenSize)
	scanner.Scan()

	in := strings.Split(scanner.Text(), " ")
	//maxProducts, _ := strconv.Atoi(in[0])
	c, _ := strconv.Atoi(in[1])

	scanner.Scan()
	sub := strings.Split(scanner.Text(), " ")

	var sums []int64
	for _, s := range sub {
		ss, _ := strconv.ParseInt(s, 10, 64)
		sums = append(sums, ss)
	}

	sort.Slice(sums, func(i, j int) bool {
		return sums[i] > sums[j]
	})

	var sum int64
	for i, v := range sums {
		sum += v
		sums[i] = sum
	}

	for i := 0; i < c; i++ {
		scanner.Scan()
		sub := strings.Split(scanner.Text(), " ")

		x, _ := strconv.Atoi(sub[0])
		y, _ := strconv.Atoi(sub[1])

		if x == y {
			fmt.Println(sums[x - 1])
		} else {
			fmt.Println(sums[x-1] - sums[x-y-1])
		}
	}
}
