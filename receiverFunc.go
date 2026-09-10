package main

import "fmt"

type User struct {
	Name string // member variable or property
	Age  int
}

// Receiver function / Method
func (u User) printUser() {
	fmt.Println("Name :", u.Name)
	fmt.Println("Age :", u.Age)
}

func main() {
	var user1 User
	user1 = User{
		Name: "John",
		Age:  30,
	}

	fmt.Println(user1)
	fmt.Println(user1.Name)
	fmt.Println(user1.Age)

	user2 := User{
		Name: "siam",
		Age:  232,
	}

	user2.printUser()
}