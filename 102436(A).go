package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var inputs []string
	var a, b, x, result, v, rem int
	var isBreak bool

	for i := 0; i < 3; i++ {
		scanner.Scan()
		inputs = append(inputs, scanner.Text())
	}

	result = 0

	a, _ = strconv.Atoi(inputs[0])
	b, _ = strconv.Atoi(inputs[1])
	x, _ = strconv.Atoi(inputs[2])

	if x == 100 {
		result = (1000 / a) * a
	} else if x == 0 {
		result = (1000 / b) * b
	} else {
		for i := 1; i < 1001; i++ {
			for j := 1; j < 1001; j++ {
				mn := a*i + b*j
				v, rem = (100*a*i)/mn, (100*a*i)%mn

				if x == v && rem == 0 {
					result = (1000 / mn) * mn
					isBreak = true
					break
				}
			}

			if isBreak {
				break
			}
		}
	}

	fmt.Println(result)
}
