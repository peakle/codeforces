package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxTokenSize = 262144 * 10

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, maxTokenSize), maxTokenSize)
	scanner.Scan()

	s := strings.Split(scanner.Text(), " ")
	w, _ := strconv.ParseInt(s[0], 10, 64)
	h, _ := strconv.ParseInt(s[1], 10, 64)
	k, _ := strconv.ParseInt(s[2], 10, 64)

	var res int64
	for i := 0; int64(i) < k; i++ {
		res += (w * h) - ((w-2)*(h-2))

		w -= 4
		h -= 4
	}

	fmt.Println(res)
}
