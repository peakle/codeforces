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
	input := scanner.Text()

	testCount, _ := strconv.Atoi(input)
	dayMin := 1440

	var timeList []string
	var hour, min, result int

	for i := 0; i < testCount; i++ {
		scanner.Scan()
		input = scanner.Text()

		timeList = strings.Split(input, " ")

		hour, _ = strconv.Atoi(timeList[0])
		min, _ = strconv.Atoi(timeList[1])

		result = dayMin - (hour*60 + min)

		fmt.Println(result)
	}
}
