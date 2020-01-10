package main

import (
    "fmt"
    "time"
)

func foo() {
    for i := 0; i < 5; i++ {
        fmt.Println("loop in foo:", i)
        time.Sleep(1 * time.Second)
    }
}

func main() {
    go foo()

    for i := 0; i < 5; i++ {
        fmt.Println("loop in main:", i)
        time.Sleep(1 * time.Second)
    }
    time.Sleep(6 * time.Second)
}