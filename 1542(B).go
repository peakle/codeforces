package main

import (
	"bufio"
	"fmt"
	"os"
)

// если не проходят базовые случаи
// i - номер случая
// посчитать сумму геометрической прогрессии n / (i * a) элемента, проверить входит ли в арифмитическую прогрессию
// посчитать сумму арифимитической прогрессии n - (i * b)
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	var testCount int
	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ {
		var n, a, b int64
		fmt.Fscan(in, &n, &a, &b)

		var yes bool
		var sumG int64 = 1
		for ; sumG <= n; sumG *= a {
			if (n-sumG)%b == 0 {
				yes = true
				break
			}
			if a == 1 {
				break
			}
		}

		if yes {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}

	_ = out.Flush()
}
