package main

import (
	"fmt"
)

// Требуется найти в бинарном векторе самую длинную последовательность единиц и вывести её длину.
// Желательно получить решение, работающее за линейное время и при этом проходящее по входному массиву только один раз.
func main() {
	count := 0

	fmt.Scan(&count)

	var s string
	max := 0
	current := 0
	for i := 0; i < count; i++ {
		fmt.Scan(&s)

		if s == "0" {
			if current > max {
				max = current
			}
			current = 0
			continue
		}

		current++
	}

	if current > max {
		max = current
	}

	fmt.Println(max)
}
