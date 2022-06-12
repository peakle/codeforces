package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const maxTokenSize = 262144 * 10

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 0, maxTokenSize), maxTokenSize)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	body := scanner.Text()
	var res int64

	fmt.Println(res % 998244353)
}
