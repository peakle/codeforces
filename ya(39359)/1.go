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

    d, _ := strconv.ParseInt(scanner.Text(), 10, 64)

    if d == 1 {
        fmt.Println(1)
        return
    }

    var sum int64 = 1
    var n int64 = 2
    var i int64 = 1
    for {
        sum += n

        if sum > d {
            break
        }

        i++
        n++
    }

    fmt.Println(i)
}
