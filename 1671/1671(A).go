package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    count, _ := strconv.Atoi(scanner.Text())

out:
    for i := 0; i < count; i++ {
        scanner.Scan()
        text := scanner.Text()

        var cos int
        var prev rune
        for ii, r := range text {
            if ii == 0 || r == prev {
                prev = r
                cos++
                continue
            }

            if r != prev {
                if cos == 1 {
                    fmt.Println("NO")
                    continue out
                }

                cos = 1
                prev = r
            }
        }

        if cos == 1 {
            fmt.Println("NO")
        } else {
            fmt.Println("YES")
        }
    }
}
