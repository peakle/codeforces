package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

const (
    lockState = "lock"
    unlockState = "unlock"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()

    color := "blue"
    lock := false

    cmdNum, _ := strconv.Atoi(scanner.Text())

    for i := 0 ; i < cmdNum ; i ++ {
        scanner.Scan()
        cmd := scanner.Text()

        switch cmd {
        case lockState:
            lock = true
        case unlockState:
            lock = false
        default:
            if !lock {
                color = cmd
            }
        }
    }

    fmt.Println(color)
}
