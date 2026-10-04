package main

import (
	"fmt"
)

var a = 23

func main() {
	x := 34

	if x > 30 {
	//	a = 100  // shadowing the outer variable 'a' with a new value
		a := 100 // shadowing the outer variable 'a' with a new variable 'a' in this block
		fmt.Println("The value of a is:", a)
	}

	fmt.Println("The value of a is:", a)
}
