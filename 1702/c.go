package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func main() {
	defer out.Flush()

	var testNum int

	fmt.Fscan(in, &testNum)
	for i := 0; i < testNum; i++ {
		n, k := 1, 1
		fmt.Fscan(in, &n, &k)

		stations := make(map[int][]int)
		for ii := 0; ii < n; ii++ {
			x := 1
			fmt.Fscan(in, &x)
			stations[x] = append(stations[x], ii)
		}

		for ii := 0; ii < k; ii++ {
			a, b := 1, 1
			fmt.Fscan(in, &a, &b)

			if _, ok := stations[a]; !ok {
				fmt.Fprintln(out, "no")
				continue
			}

			if _, ok := stations[b]; !ok {
				fmt.Fprintln(out, "no")
				continue
			}

			v1 := stations[a][0]
			v2 := stations[b][len(stations[b])-1]

			if v1 < v2 {
				fmt.Fprintln(out, "yes")
			} else {
				fmt.Fprintln(out, "no")
			}
		}
	}
}
