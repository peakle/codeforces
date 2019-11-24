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
	var content string
	var a, b, x, result, v, rem int
	var isBreak bool

	for scanner.Scan() {
		result = 0
		content = scanner.Text()
		inputs := strings.Split(content, " ")
		a, _ = strconv.Atoi(inputs[0])
		b, _ = strconv.Atoi(inputs[1])
		x, _ = strconv.Atoi(inputs[2])

	if x == 100 {
		result = (1000 / a) * a
	} else if x == 0 {
		result = (1000 / b) * b
	} else {
		for i := 1; i < 1000; i++ {
			for j := 1; j < 1000; j++ {
				mn := a*i + b*j
				v, rem = (100*a*i)/mn, (100*a*i)%mn

				if x == v && rem == 0 {
					result = (1000 / mn) * mn
					fmt.Println(result)
					isBreak = true
					break
				}
			}

			if isBreak {
				isBreak = false
				break
			}
		}
	}
	}
}
