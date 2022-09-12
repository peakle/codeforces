package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

const maxTokenSize = 2500000 * 7

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 0, maxTokenSize), maxTokenSize)
    scanner.Scan()

    count, _ := strconv.Atoi(scanner.Text())

out:
    for i := 0; i < count; i++ {
        scanner.Scan()
        dotsCount, _ := strconv.Atoi(scanner.Text())

        var dots []int
        scanner.Scan()
        for _, dot := range strings.Split(scanner.Text(), " ") {
            d, _ := strconv.Atoi(dot)
            dots = append(dots, d)
        }

        if dotsCount == 1 {
            fmt.Println("YES")
            continue
        }

        var diff int
        for i := range dots {
            if len(dots)-1 == i {
				break
            }

            diff += dots[i+1] - dots[i] - 1
        }

        if diff <= 2 {
            fmt.Println("YES")
            continue out
        }

        fmt.Println("NO")
    }
}
