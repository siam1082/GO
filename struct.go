package main 
 
import "fmt"


type User struct {
	Name string // member varibale or property 
	Age int 

}



func main () {
	var user1 User 
	user1 = User{ // instantiate 
		Name: "John", Age: 30,
	}
	fmt.Println(user1)
	fmt.Println(user1.Name)
	fmt.Println(user1.Age)

	user2 := User{ // instance of an object 
		Name : "siam",
		Age : 232,
	}

	fmt.Println(user2)
	fmt.Println(user2.Name)
	fmt.Println(user2.Age)
 }

