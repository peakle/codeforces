package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	var testCount int
	var n, r int64
	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ {
		fmt.Fscan(in, &n)

		list := make([]int64, 0, n)
		for k := int64(0); k < n; k++ {
			fmt.Fscan(in, &r)
			list = append(list, r)
		}

		var sum int64
		for _, k := range list {
			sum += k
		}
		base := sum / n
		dif := sum % n

		for ii := range list {
			if dif > 0 {
				list[ii] = base + 1
				dif--
				continue
			}
			list[ii] = base
		}

		fmt.Fprintln(out, calc(list))
	}

	_ = out.Flush()
}

func calc(start []int64) int64 {
	var temp []int64

	for i := 0; i < len(start); i++ {
		for j := i + 1; j < len(start); j++ {
			d := start[i] - start[j]
			if d < 0 {
				d *= -1
			}

			temp = append(temp, d)
		}
	}

	var res int64 = 0
	for i := 0; i < len(temp); i++ {
		res += temp[i]
	}

	return res
}
