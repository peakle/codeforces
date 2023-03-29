package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 5 1
// 0 0
// 0 2
// 2 0
// 1 1
// 2 2

// 2
// 1 4

// 5 2
// 0 0
// 0 2
// 2 0
// 1 1
// 2 2

// 3
// 1 3 4

// 5 4
// 0 0
// 0 2
// 2 0
// 1 1
// 2 2

// 5
// 1 2 3 4 5

// 2 1
// 0 0
// 2 0

// 1
// 1

type Pos struct {
	X, Y, i int64
}

func main() {
	var n int
	var s int64
	fmt.Scan(&n, &s)

	var orders []Pos
	reader := bufio.NewReader(os.Stdin)
	index := int64(1)

	for i := 0; i < n; i++ {
		line, _, _ := reader.ReadLine()

		xy := strings.Split(string(line), " ")

		x, _ := strconv.ParseInt(xy[0], 10, 64)
		y, _ := strconv.ParseInt(xy[1], 10, 64)

		orders = append(orders, Pos{
			X: x,
			Y: y,
			i: index,
		})

		index++
	}

	res, s := searchRect(orders, s)

	fmt.Println(len(res))
	for _, r := range res {
		fmt.Println(r.i)
	}
}

func searchRect(points []Pos, s int64) ([]Pos, int64) {
	sort.Slice(points, func(i, j int) bool {
		return points[i].X < points[j].X
	})

	return nil, 0
}
