package main

import "fmt"

var {
	arr2 = [3]string{"apple", "banana", "cherry"}
}

func main () {

	var arr [5]int 

	fmt.Println(arr)

	arr1 := [5]int{1,2,3,4,5}
	fmt.Println(arr1)

	arr[1] = 9
	fmt.Println(arr)
	  
}

