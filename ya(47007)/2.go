package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	var counter int
	fmt.Scan(&counter)

	res := []float64{0, 0, 0, 0}
	for i := 0; i < counter; i++ {
		var sum float64
		var start, end string
		fmt.Scan(&sum, &start, &end)

		startT, _ := time.Parse("02.01.2006", start+".2022")
		endT, _ := time.Parse("02.01.2006", end+".2022")

		days := (endT.Sub(startT).Hours() / 24) + 1

		perDay := math.Floor((sum / days) * 100)

		for {
			// (1-1) /3 = 0
			// (2-1) /3 = 0
			// ...
			// (11-1) /3 = 3
			// (12-1) /3 = 3
			quart := (int(startT.Month()) - 1) / 3
			res[quart] += perDay
			startT = startT.Add(24 * time.Hour)

			if startT.After(endT) && !startT.Equal(endT) {
				break
			}
		}
	}

	for _, r := range res {
		fmt.Printf("%.2f\n", r/100)
	}
}
