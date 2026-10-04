package main

import "fmt"

func a(){
	i :=0
	defer fmt.Println("defer :", i)

	i++
	fmt.Println(i)
	
	defer fmt.Println("defer2 :", i)  // --> last in first out 



	return

}
func calculate() (result int ){
	fmt.Println("first :", result)

	show := func(){
		result= result+5
		fmt.Println("defer :", result)
	}

	defer show()

	result = 5

	fmt.Println("second :", result )

	return 

}


func calc() int {
	result := 0
	fmt.Println("first :", result)

	show := func(){
		result= result+5
		fmt.Println("defer :", result)
	}

	defer show()

	result = 5

	fmt.Println("second :", result )

	return result

}
func main(){
	//a()
	a := calculate()
	fmt.Println("main first :",a)


	b:=calc()
	fmt.Println("main second : ",b)
}




