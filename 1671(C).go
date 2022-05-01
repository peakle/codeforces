package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const maxTokenSize = 25000000 * 7

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 0, maxTokenSize), maxTokenSize)
	scanner.Scan()

	countSum := func(pieces []int64) []int64 {
		var sum = make([]int64, 0, len(pieces))
		var local int64
		for _, piece := range pieces {
			local += piece
			sum = append(sum, local)
		}

		return sum
	}

	count, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < count; i++ {
		scanner.Scan()
		s := strings.Split(scanner.Text(), " ")

		shopCount, _ := strconv.ParseInt(s[0], 10, 64)
		moneyForDay, _ := strconv.ParseInt(s[1], 10, 64)

		prices := make([]int64, 0, shopCount)
		scanner.Scan()

		for _, p := range strings.Split(scanner.Text(), " ") {
			price, _ := strconv.ParseInt(p, 10, 64)
			prices = append(prices, price)
		}

		sort.Slice(prices, func(i, j int) bool {
			return prices[i] < prices[j]
		})

		var dayNum int64
		var sugarSum int64
		sum := countSum(prices)

		for {
			startSum := moneyForDay - (sum[shopCount-1] + shopCount*dayNum)
			if startSum < 0 {
				shopCount--
				if shopCount == 0 {
					break
				}

				continue
			}

			days := (startSum / shopCount) + 1

			dayNum += days
			sugarSum += shopCount * days
		}

		fmt.Println(sugarSum)
	}
}
