package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const firstLetter = 'a' - 1

//1
//zzcd
//52
func main() {
	var in = bufio.NewReader(os.Stdin)

	var testNum int
	fmt.Fscan(in, &testNum)

	for ; testNum > 0; testNum-- {
		var s string
		fmt.Fscan(in, &s)

		var p int
		fmt.Fscan(in, &p)

		var w int
		for _, l := range s {
			v := int(l - firstLetter)
			w += v
		}

		d := p - w
		if d >= 0 {
			fmt.Println(s)
			continue
		}

		ss := []rune(s)
		sort.Slice(ss, func(i, j int) bool {
			return ss[i] > ss[j]
		})

		del := make(map[rune]int, 10)
		for _, sss := range ss {
			if w > p {
				w -= int(sss - firstLetter)
				del[sss]++
			} else {
				break
			}
		}

		se := []rune(s)
		for i, v := range se {
			if vv, ok := del[v]; ok && vv > 0 {
				del[v]--
				se[i] = 0
			}
		}

		res := make([]rune, 0, len(se))
		for _, ss := range se {
			if ss != 0 {
				res = append(res, ss)
			}
		}

		fmt.Println(string(res))
	}
}
