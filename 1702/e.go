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

		dominos := make([]domin, 0, n)

		var failed bool
		for ; n > 0; n-- {
			a, b := 0, 0
			fmt.Fscan(in, &a, &b)

			if a == b {
				failed = true
			}
			dominos = append(dominos, domin{
				a: a,
				b: b,
			})
		}

		if failed {
			fmt.Println("no")
			continue
		}

		fg, sg := make([]domin, 0, len(dominos)/2), make([]domin, 0, len(dominos)/2)
		empty := domin{}

		uniq := map[int]struct{}{}
		for i, d := range dominos {
			if d == empty {
				continue
			}

			_, ok := uniq[d.a]
			_, ok2 := uniq[d.b]
			if !ok && !ok2 {
				uniq[d.a] = struct{}{}
				uniq[d.b] = struct{}{}
				dominos[i] = empty
				fg = append(fg, domin{
					a: d.a,
					b: d.b,
				})
			}
		}

		uniq = map[int]struct{}{}
		for i, d := range dominos {
			if d == empty {
				continue
			}

			_, ok := uniq[d.a]
			_, ok2 := uniq[d.b]
			if !ok && !ok2 {
				uniq[d.a] = struct{}{}
				uniq[d.b] = struct{}{}
				dominos[i] = empty
				sg = append(sg, domin{
					a: d.a,
					b: d.b,
				})
			} else {
				break
			}
		}

		if len(sg)+len(fg) == len(dominos) {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
