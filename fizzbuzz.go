package main

import "fmt"

func main() {
    for i := 1; i <= 100; i++ {

        if i%3 == 0 {
            fmt.Printf("fizz")
        }
        if i%5 == 0 {
            fmt.Printf("buzz")
        }

        if i%3 != 0 && i%5 != 0 {
            fmt.Printf("%d", i)
        }
        fmt.Printf("\n")

    }
}
