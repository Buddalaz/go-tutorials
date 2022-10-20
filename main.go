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
	// nameFour := "four"
	// fmt.Println(nameFour)

	//int
	var ageOne int = 20
	var ageTwo = 30
	//infort type
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	//bits & memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numeThree uint16 = 256

	var scoreOne float32 = 25.98
	var scoreTwo float64 = 3454534434234424325.98
	scoreThree := 1.5
}
