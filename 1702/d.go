package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in = bufio.NewReader(os.Stdin)

	var testNum int
	fmt.Fscan(in, &testNum)

	for ; testNum > 0; testNum-- {
		var s string
		fmt.Fscan(in, &s )

		var p int64
		fmt.Fscan(in, &p)


	}
}
