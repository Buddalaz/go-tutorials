package main

import "fmt"

func main() {

	//create/decalre/initialize a variables in go
	//strings

	//1.data type definition
	var nameOne string = "one"
	fmt.Println(nameOne)

	//2.without data type
	var nameTwo = "two"
	fmt.Println(nameTwo)

	//3.assign a variable later
	var nameThree string
	fmt.Println(nameThree)

	nameThree = "three" //can't change data type like assign int value to varaible nameThree
	fmt.Println(nameThree)

	//4.without using var and use instead colun we only do this when initializing declaring the variable
	//cant use this outside a function
	nameFour := "four"
	fmt.Println(nameFour)

}
