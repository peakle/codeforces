package main

import (
	"os"
	"bufio"
	"strconv"
	"fmt"
)

func main()  {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())

	var num int
	for i := 0 ; i < count; i++ {
		scanner.Scan()
		num ,  _ = strconv.Atoi(scanner.Text())

		if num %4 == 0 {
			fmt.Println("YES")

			continue
		}

		fmt.Println("NO")
	}
}
