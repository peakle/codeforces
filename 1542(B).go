package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// если не проходят базовые случаи
// i - номер случая
// посчитать сумму геометрической прогрессии n / (i * a) элемента, проверить входит ли в арифмитическую прогрессию
// посчитать сумму арифимитической прогрессии n - (i * b)
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	testCount, _ := strconv.Atoi(scanner.Text())

	inputList := make([]int, 0, 3)
next:
	for i := 0; i < testCount; i++ {
		inputList = inputList[:0]

		scanner.Scan()
		input := scanner.Text()

		for _, p := range strings.Split(input, " ") {
			o, _ := strconv.Atoi(p)
			inputList = append(inputList, o)
		}

		n, a, b := inputList[0], inputList[1], inputList[2]

		if n%a == 0 && a != 1 {
			fmt.Println("Yes")
			continue
		}

		if n%b == 1 || b == 1 {
			fmt.Println("Yes")
			continue
		}

		if a == 1 {
			fmt.Println("No")
			continue
		}

		sumG := 1
		prevSum := 1
		for sumG < n {
			prevSum = sumG
			sumG = prevSum * a

			if (n-sumG)%b == 0 {
				fmt.Println("Yes")
				continue next
			}
		}

		fmt.Println("No")
	}
}
