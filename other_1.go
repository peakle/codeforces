package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/**
Задача C. Полина составляет программу [C]
Имя входного файла: стандартный ввод
Имя выходного файла: стандартный вывод
Ограничение по времени: 1 секунда
Ограничение по памяти: 256 мегабайт

Полина решила составить план занятий параллели C на следующий год. У нее на примете есть
задачи на n разных тем, причем на i-ю тему у нее есть ai задач.
Будем называть программу интересной, если среди любых k подряд идущих задач нет пары на одинаковую тему.
Разумеется, чем больше задач, тем лучше программа. Помогите Полине посчитать максимальное
количество задач, вошедших в интересную программу.

Формат входных данных
В первой строке находятся числа n и k (1 6 n, k 6 105) — количество тем и параметр k.
Следующие n строк содержат числа ai (1 6 ai 6 109) (по одному в строке) — количество задач на i-ю тему.

Формат выходных данных:
Выведите единственное число — максимальное число задач в интересной программе.

Пример

стандартный ввод:
2 2
10
50

стандартный вывод:
21
 */

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()

	inputs := strings.Split(input, " ")

	n, _ := strconv.Atoi(inputs[0])
	k, _ := strconv.Atoi(inputs[1])

	if k > n {
		fmt.Println(n)
		return
	}

	arr := make([]int, n)

	var element,index int
	for index = 0; index < n; index++ {
		scanner.Scan()
		element, _ = strconv.Atoi(scanner.Text())
		arr[index] = element
	}

	sort.Ints(arr)
	var res, prev int

	arrLen := len(arr)
	lastIndex := arrLen - k

	for index, element = range arr {
		res += (element - prev) * arrLen

		if index == lastIndex {
			res++
			break
		}

		arrLen--
		prev = element
	}

	fmt.Println(res)
}
