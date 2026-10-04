package main

import (
	"fmt"
)

func processData (a int , b int , op func(x int, y int) int) int {
	c := op(a, b)

	return op(a, b)

}