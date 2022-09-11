package main

import (
    "bufio"
    "fmt"
    "os"
)

//not finished
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()

    text := []rune(scanner.Text())

    p := 0
    dup := 0
    var prev rune
    for i, l := range text {
        if l == prev {
            dup++
        } else {
            dup = 0
        }

        switch {
        case '/' == l:
            if dup == 0 {
                text[p] = text[i]
                p++
            }
        case '.' == l:
            if dup == 2 {
                if len(text) > i+2 {
                    if text[i+2] != '.' {
                        p--
                    }
                    //p++
                    //} else {
                    //if len(text) > 3 {
                    //	p -= 2
                    //}
                }
            }
        default:
            text[p] = text[i]
            p++
        }

        prev = l
    }

    if p > 0  && text[p-1] == '/' && len(text[:p]) > 1 {
		p--
    }

    fmt.Println(text[:p])
}
