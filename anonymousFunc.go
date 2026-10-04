package main

import (
	"fmt"
)


func init() {

	fmt.Println("init function executed")

}

func main (){
	func (a int , b int) {

	c:= a + b
	fmt.Println("The result is:", c)
	}(23, 34)  // here we immediately invoke the anonymous function with arguments 23 and 34
	
}
