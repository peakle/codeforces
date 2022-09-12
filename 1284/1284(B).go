package main

import (
    "bufio"
"fmt"
"os"
"sort"
"strconv"
"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer([]byte{}, 1000001000)
	var result uint64
	scanner.Scan()
	var n, ind int
	n, _ = strconv.Atoi(scanner.Text())

	var maxv []int
	var minv []int
	var sCount, l int

	var sList []string
	for ; sCount < n; sCount++ {
		scanner.Scan()

		sList = strings.Split(scanner.Text(), " ")
		l, _ = strconv.Atoi(sList[0])

		s := make([]int, l)

		for ind = 0; ind < l; ind++ {
			s[ind], _ = strconv.Atoi(sList[ind+1])
		}

		if l == 1 || sort.IsSorted(sort.Reverse(sort.IntSlice(s))) {
			sort.Ints(s)
			maxv = append(maxv, s[l-1])
			minv = append(minv, s[0])
		} else {
			maxv = append(maxv, 1e9)
			minv = append(minv, -1e9)
		}
	}

	sort.Ints(maxv)

	var i int

	for i = 0; i < len(maxv); i++ {
		result += uint64(n - sort.Search(len(maxv), func(j int) bool {
			return maxv[j] > minv[i]
		}))
	}

	fmt.Println(result)
}
