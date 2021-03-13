package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type setType struct {
	s1, s2 int
}

var ar []int

const maxTokenSize = 2500000 * 7

func main() {
	var size int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 0, maxTokenSize), maxTokenSize)

	scanner.Scan()

	size, _ = strconv.Atoi(scanner.Text())

	var ar = make([]int, 0, size)
	var sums = make(map[int]setType, size)
	var val int

	scanner.Scan()

	inputs := strings.Split(scanner.Text(), " ")

	for _, v := range inputs {
		val, _ = strconv.Atoi(v)
		ar = append(ar, val)
	}

	for s1, v1 := range ar {
		for e1, v2 := range ar {
			if s1 >= e1 {
				continue
			}

			sum := v1 + v2
			if set, ok := sums[sum]; ok && s1 != set.s1 && s1 != set.s2 && e1 != set.s1 && e1 != set.s2 {
				fmt.Println("YES")
				fmt.Println(s1+1, e1+1, set.s1+1, set.s2+1)
				return
			} else if ok == false {
				sums[sum] = setType{s1: s1, s2: e1}
			}
		}
	}

	fmt.Println("NO")
}
