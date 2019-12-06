package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var result int
	var content string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	content = scanner.Text()
	n, _ := strconv.Atoi(content)

	for i := 0; i < n; i++ {
		content = scanner.Text()
		inputs := strings.Split(content, " ")

		c, _ := strconv.Atoi(inputs[0])
		sum, _ := strconv.Atoi(inputs[1])

		result = calcMinSum(c, sum)
		fmt.Println(result)
	}
}

func calcMinSum(c, sum int) int {


	return 0
}
