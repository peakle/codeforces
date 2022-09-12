package main

import (
    "bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReaderSize(os.Stdin, 10000)

	var testCount int
	var n, r int64
	fmt.Fscan(in, &testCount)
	list := make([]int64, 0, n)

	for i := 0; i < testCount; i++ {
		fmt.Fscan(in, &n)

		list = list[:0]
		for k := int64(0); k < n; k++ {
			fmt.Fscan(in, &r)
			list = append(list, r)
		}

		var sum int64
		for _, k := range list {
			sum += k
		}
		dif := sum % n

		fmt.Println((int64(len(list)) - dif) * dif)
	}
}
