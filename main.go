package main

import "fmt"

func main() {

	//Loops
	//for all the loops ex:- while, for, for in in go use for
	x := 0
	for x < 5 {
		fmt.Println("The value of x is:", x)
		x++
	}
	//output will be
	// The value of x is: 0
	// The value of x is: 1
	// The value of x is: 2
	// The value of x is: 3
	// The value of x is: 4

	//traditional loop
	for i := 0; i < 5; i++ {
		fmt.Println("The value of i is:", i)
	}
	//output will be
	// The value of x is: 0
	// The value of x is: 1
	// The value of x is: 2
	// The value of x is: 3
	// The value of x is: 4

	//loop throught slice or string or numbers
	names := []string{"animal", "bird", "fish", "reptails"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	//output will be
	// animal
	// bird
	// fish
	// reptails

	//another way to run slice just like for in instead we defined the index and value and define ':= range'
	for index, value := range names {
		fmt.Printf("The value at index %v is %v \n", index, value)
	}
	//output will be
	// 	The value at index 0 is animal
	// The value at index 1 is bird
	// The value at index 2 is fish
	// The value at index 3 is reptails

	//if we don't want to used index we can just leave it like this '_'
	for _, value := range names {
		fmt.Printf("The value is %v \n", value)
	}
	//output will be
	// 	The value is animal
	// The value is bird
	// The value is fish
	// The value is reptails

	for _, value := range names { //value is local copy of a variable
		fmt.Printf("The value is %v \n", value)
		value = "new string" //this value does not update the original slice
	}

	fmt.Println(names)
}
