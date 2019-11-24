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
	var bars []string
	var length float64
	var err error
	var positionA, positionB float64

	for scanner.Scan() {
		content = scanner.Text()
		bars = strings.Split(content, " ")
		positionA, err = strconv.ParseFloat(bars[0], 0)
		if err != nil {
			return
		}

		positionB, err = strconv.ParseFloat(bars[1], 0)
		if err != nil {
			return
		}

		length, err = strconv.ParseFloat(bars[2], 0)
		if err != nil {
			return
		}

		diffLen := positionB - positionA
		count := 0

		for {
			if positionB >= length {
				if positionA >= length {
					break
				}

				positionA = positionB
				count++
				break
			}

			if positionA >= length {
				if positionB >= length {
					break
				}

				positionB = positionA
				count++
				break
			}

			if positionA > positionB {
				positionB = positionA
				count++
				continue
			}

			if positionB > positionA {
				positionA = positionB
				count++
				continue
			}

			if positionA == positionB {
				positionB += diffLen
				count++
				continue
			}
		}

		fmt.Println(count)
	}
}
