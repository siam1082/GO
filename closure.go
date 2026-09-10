package main

import "fmt"

const a = 10

var p = 100

func outer() func(){
	money := 100
	age :=300
	fmt.Println("AGE" , age)
	show := func() {
		money = money + a + p
		fmt.Println(money)
	}
	return show
}


func call(){
	incr1 := outer()
	incr1()
	incr1()

	incr2 := outer()
	incr2()
	incr2()	
}
func main() {
	call()
}

func init(){
	fmt.Println("=== Bank ===")
}





// compilation phase 
// executaion phase 

// go build closure.go -> compile 
// ./main -> execute

// ========================
// code segmnent

// a = 100 
// outer= func () ..
// outerAnonymous1 =  func ()
// call = gunc ()..()
// main = func ().
// init = func ()... 



