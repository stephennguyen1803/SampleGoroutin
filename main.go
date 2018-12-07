package main

import (
    "fmt"
    "log"
    "time"
)

func checkNumberWithGoRoutin(n int, c chan bool) bool {
    var fib1, fib2, fib = 1, 1, 2

    for  fib1 + fib2 <= n {
        fib = fib1 + fib2
        fib2 = fib1
        fib1 = fib
    }
    c <- fib == n
    return fib == n
}


func checkNumberNormal(n int) bool {
    var fib1, fib2, fib = 1, 1, 2

    for  fib1 + fib2 <= n {
        fib = fib1 + fib2
        fib2 = fib1
        fib1 = fib
    }

    return fib == n
}

func main() {

    start := time.Now()

    numbers := []int{4, 36 , 16, 25, 85, 91, 40, 36, 88, 5}

    c := make(chan bool)

    for _, number := range numbers {

        go checkNumberWithGoRoutin(number, c)

        if ( <- c) {
            fmt.Println(number,"co so fibonace")
            elapsed := time.Since(start)
            log.Printf("using goroutin took %s", elapsed)
        }
    }
}
