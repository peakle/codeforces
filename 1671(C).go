package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const maxTokenSize = 4 * 100000

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 0, maxTokenSize), maxTokenSize)
	scanner.Scan()

	countSum := func(pieces []int) int {
		var sum int
		for _, piece := range pieces {
			sum += piece
		}
		return sum
	}

	count, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < count; i++ {
		scanner.Scan()
		s := strings.Split(scanner.Text(), " ")

		shopCount, _ := strconv.Atoi(s[0])
		moneyForDay, _ := strconv.Atoi(s[1])

		prices := make([]int, 0, 30)
		scanner.Scan()

		for _, p := range strings.Split(scanner.Text(), " ") {
			price, _ := strconv.Atoi(p)
			prices = append(prices, price)
		}

		sort.Ints(prices)

		var dayNum int
		var sugarSum int
		for {
			sum := countSum(prices) + shopCount*dayNum
			if sum > moneyForDay  {
				shopCount--
				if shopCount == 0 {
					break
				}

				prices = prices[:shopCount]
				continue
			}

			dayNum++
			sugarSum += shopCount
		}

		fmt.Println(sugarSum)
	}
}
