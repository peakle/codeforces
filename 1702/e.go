package main

import (
	"bufio"
	"fmt"
	"os"
)

type domin struct {
	a, b int
}

// https://neerc.ifmo.ru/wiki/index.php?title=%D0%A1%D0%9D%D0%9C_(%D0%BD%D0%B0%D0%B8%D0%B2%D0%BD%D1%8B%D0%B5_%D1%80%D0%B5%D0%B0%D0%BB%D0%B8%D0%B7%D0%B0%D1%86%D0%B8%D0%B8)
func main() {
	in := bufio.NewReader(os.Stdin)

	var testNum int
	fmt.Fscan(in, &testNum)

	for ; testNum > 0; testNum-- {
		n := 0
		fmt.Fscan(in, &n)

		dominos := make(map[int][]int, n)

		var failed bool
		for N := n; N > 0; N-- {
			a, b := 0, 0
			fmt.Fscan(in, &a, &b)

			dominos[a] = append(dominos[a], b)
			dominos[b] = append(dominos[b], a)

			if a == b || len(dominos[a]) > 2 || len(dominos[b]) > 2 {
				failed = true
			}
		}

		if failed {
			fmt.Println("no")
			continue
		}

		used := make(map[int]bool)
		var fn func(v int) int

		fn = func(v int) int {
			used[v] = true
			for _, now := range dominos[v] {
				if !used[now] {
					return fn(now) + 1
				}
			}
			return 1
		}

		for i := 0; i < n; i++ {
			if !used[i+1] {
				if fn(i+1)%2 > 0 {
					failed = true
					break
				}
			}
		}

		if failed {
			fmt.Println("no")
		} else {
			fmt.Println("yes")
		}
	}
}
