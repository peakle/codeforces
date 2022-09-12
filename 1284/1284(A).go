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

	inputList := strings.Split(scanner.Text(), " ")

	n, _ := strconv.Atoi(inputList[0])
	m, _ := strconv.Atoi(inputList[1])

	nList := make(map[int]string, n+1)
	mList := make(map[int]string, m+1)

	var index int
	var element string

	scanner.Scan()
	for index, element = range strings.Split(scanner.Text(), " ") {
		nList[index+1] = element
	}

	scanner.Scan()
	for index, element = range strings.Split(scanner.Text(), " ") {
		mList[index+1] = element
	}

	scanner.Scan()
	q, _ := strconv.Atoi(scanner.Text())

	var quest, remN, remM int
	for i := 0; i < q; i++ {
		scanner.Scan()
		quest, _ = strconv.Atoi(scanner.Text())

		remN = quest % n
		if remN == 0 {
			remN = n
		}

		remM = quest % m
		if remM == 0 {
			remM = m
		}

		fmt.Println(nList[remN] + mList[remM])
	}
}
