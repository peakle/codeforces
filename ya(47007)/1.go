package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var w float64
	var counter, k int
	fmt.Scan(&w, &counter, &k)
	var hMax, hMin float64 = 0, 1e18

	for i := 0; i < counter; i++ {
		var hw string
		fmt.Scan(&hw)

		s := strings.Split(hw, "x")

		wi, _ := strconv.ParseFloat(s[0], 64)
		hi, _ := strconv.ParseFloat(s[1], 64)

		h := math.Ceil(hi * (w / wi))

		if h > hMax {
			hMax = h
		}
		if h < hMin {
			hMin = h
		}
	}

	fmt.Println(hMax)
	fmt.Println(hMin)
}
