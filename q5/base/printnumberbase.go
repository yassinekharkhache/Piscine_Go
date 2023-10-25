package main

import "fmt"


func PrintNbrBase(num int, base string) {
    if len(base) < 2 {
        fmt.Println("NV")
        return
    }

    if len(base) > len([]byte(base)) {
        fmt.Println("NV")
        return
    }

    result := ""
    negative := false
    if num < 0 {
        negative = true
        num = -num
    }

    for num > 0 {
        result = string( base[num % len(base) ]) + result
        num = num / len(base)
    }

    if negative {
        result = "-" + result
    }

    fmt.Println(result)
}
