package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// Pass by value
func changeArray1(arr [3]int) {
	arr[0] = 100
}

// Pass by reference
func changeArray2(arr *[3]int) {
	arr[0] = 200
}

func main() {
	// Pointer or memory address
	x := 20

	fmt.Println(x)

	p := &x

	fmt.Println(p)  // address
	fmt.Println(*p) // value


	// Array pointer
	arr := [3]int{10, 20, 30}

	arrPtr := &arr

	fmt.Println(arr)
	fmt.Println(arrPtr)       // address of array
	fmt.Println((*arrPtr)[0]) // first element


	// Pass by value
	changeArray1(arr)
	fmt.Println(arr) // [10 20 30]


	// Pass by reference
	changeArray2(&arr)
	fmt.Println(arr) // [200 20 30]


	// Struct pointer
	user := User{
		Name: "John",
		Age:  30,
	}

	userPtr := &user

	fmt.Println(*userPtr)
	fmt.Println(userPtr)
	fmt.Println(userPtr.Name)
	fmt.Println(userPtr.Age)
}