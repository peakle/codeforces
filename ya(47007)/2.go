package main

import (
	"fmt"
	"math"
	"time"
)

// 4
// 10 10.02 11.05
// 10 11.12 23.12
// 100 24.05 30.06
// 4342 10.07 12.09

// 5.00
// 104.04
// 4342.00
// 9.88

// 1
// 1000000 01.01 31.12

// 246574.80
// 249314.52
// 252054.24
// 252054.24
func main() {
	var counter int
	fmt.Scan(&counter)

	res := []float64{0, 0, 0, 0}
	for i := 0; i < counter; i++ {
		var sum float64
		var start, end string
		fmt.Scan(&sum, &start, &end)

		startT, _ := time.Parse("02.01.2006", start+".2022")
		endT, _ := time.Parse("02.01.2006", end+".2022") // 29 Feb will be first March

		days := (endT.Sub(startT).Hours() / 24) + 1

		perDay := math.Ceil((sum / days) * 100)

		for {
			quart := (int(startT.Month()) - 1) / 3
			res[quart] += perDay
			startT = startT.Add(24 * time.Hour)

			if startT.After(endT) && !startT.Equal(endT) {
				break
			}
		}
	}

	for _, r := range res {
		fmt.Printf("%.2f\n", float64(r/100))
	}
}
