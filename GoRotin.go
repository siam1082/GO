package main

import (
    "fmt"
    "time"
)


var a = 10
const p = 11


func printhloo(num int) {
    fmt.Println("helo Siam", num)
}

func main() {
    go printhloo(1)
    go printhloo(2)
    go printhloo(3)

	fmt.Println(a," ", p)
    time.Sleep(5 * time.Second)
}