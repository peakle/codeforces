package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	testNum, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < testNum; i++ {
		scanner.Scan()
		length, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		var nums []int
		for _, s := range strings.Split(scanner.Text(), " ") {
			n, _ := strconv.Atoi(s)
			nums = append(nums, n)
		}

		if length == 1 {
			fmt.Println(nums[0])
			continue
		}

		var (
			sum  = 1
			sumi int
			ones int
		)

		sort.Ints(nums)
		var maxOdd int

		for _, n := range nums {
			if n%2 == 0 {
				if sum != 1 {
					ones++
				}

				sum *= n
			} else {
				if n > maxOdd {
					maxOdd = n
				}

				sumi += n
			}
		}

		if maxOdd != 0 {
			sum = (sum*maxOdd + 1) + (sumi + ones - maxOdd)
			fmt.Println(sum)
			continue
		}

		if sum == 1 {
			sum = sumi
		} else {
			sum += sumi + ones
		}

		fmt.Println(sum)
	}
}

