package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"math/big"
	"os"
	"sort"
	"strconv"
)

func main() {
	var w float64
	var counter, k int
	fmt.Scan(&w, &counter, &k)

	var heights []int64
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < counter; i++ {
		hwB, _, _ := reader.ReadLine()

		s := bytes.Split(hwB, []byte("x"))

		wi, _ := strconv.ParseFloat(string(s[0]), 64)
		hi, _ := strconv.ParseFloat(string(s[1]), 64)

		h := int64(math.Ceil(hi * (w / wi)))

		heights = append(heights, h)
	}

	sort.Slice(heights, func(i, j int) bool {
		return heights[i] < heights[j]
	})

	min := big.NewInt(0)
	max := big.NewInt(0)

	for _, h := range heights[:k] {
		min = min.Add(min, big.NewInt(h))
	}
	for _, h := range heights[len(heights)-k:] {
		max = max.Add(max, big.NewInt(h))
	}

	fmt.Println(min)
	fmt.Println(max)
}
