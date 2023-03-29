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

// 2000
// 5 3
// 1000x1000
// 1000x1500
// 640x930
// 640x1500
// 3000x1000

// 5574
// 10595

// 1000
// 5 5
// 1404x1386
// 1132x1110
// 1061x1801
// 1022x1519
// 1529x1003

// 5810
// 5810

// 4096
// 2 1
// 640x4096
// 4096x640

// 640
// 26215
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
