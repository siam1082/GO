package main 

import "fmt"

func main(){
	arr := [6]string{"This", "is", "a ","go","interview", "question"}
	fmt.Println(arr)

	s := arr[1:4] // -> ptr 1, len 3, cap 5
	fmt.Println(s)

	s1 := s[1:2] // -> ptr 2, len 1, cap 4
	fmt.Println(s1)

	
}