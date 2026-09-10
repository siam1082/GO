package main

import (
	"fmt"
)

var a = 10

func main() {
	fmt.Println("main function started")
	fmt.Println("The value of a is:", a)
}

func init() {

	fmt.Println("init function executed")
	fmt.Println("The value of a is:", a)
	a = 23

}

// Package-level variables
//         ↓
//     init()
//         ↓
//     main()
