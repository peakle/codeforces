package main

import (
    "bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	c, _ := strconv.Atoi(scanner.Text())
	var inputList = make([]int, 3)
	for i := 0; i < c; i++ {
		scanner.Scan()
		for ii, d := range strings.Split(scanner.Text(), " ") {
			inputList[ii], _ = strconv.Atoi(d)
		}

		n := inputList[0]
		m := inputList[1]
		k := inputList[2]

		d := n - m
		c := d / (k - 1)
		if c < m {
			fmt.Println(c)
		} else {
			fmt.Println(m)
		}
	}
}
