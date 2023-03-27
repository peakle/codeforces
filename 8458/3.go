package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Дан упорядоченный по неубыванию массив целых 32-разрядных чисел. Требуется удалить из него все повторения.
// Желательно получить решение, которое не считывает входной файл целиком в память, т.е., использует лишь константный объем памяти в процессе работы.
func main() {
	d := make(map[string]struct{}, 100)
	res := make([]string, 0, 1000)

	var line1 string
	fmt.Scan(&line1)
	counter, _ := strconv.Atoi(line1)

	out, _ := os.Create("output.txt")

	scaneer := bufio.NewReader(os.Stdin)
	for i := 0; i < counter; i++ {
		lineB, _, _ := scaneer.ReadLine()
		line := string(lineB)

		if _, ok := d[line]; !ok {
			d[line] = struct{}{}
			res = append(res, line)
		}

		if len(res) > 1000 {
			fmt.Fprint(out, strings.Join(res, "\n")+"\n")

			res = res[:0]
		}
	}

	fmt.Fprint(out, strings.Join(res, "\n")+"\n")
}
