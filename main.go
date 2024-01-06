package main

import "fmt"

func main() {

	age := 35
	name := "sandun"

	//Print
	fmt.Print("hello, ") //start with capital letter cuz fmt makes that public, and this dosen't add a new line
	fmt.Print("world!")  // output is - hello, world!

	fmt.Print("hello, ") //if we want to add a new line use '\n'
	fmt.Print("world! \n")
	fmt.Print("new line")

	//Println
	fmt.Println("hello sandun!") //this will print separatly and add a new line
	fmt.Println("goodbye sandun!")

	// output variables to the console using Print function
	fmt.Println("my age is", age, "my name is", name) //to use a variable in Println use ',' comma to concat

	// format strings
	//is whey to create a string with variables embeded inside
	//Printf (formatted strings)
	fmt.Printf("my age is %v and my name is %v \n", age, name) //we use format specifiers to emded a variable to a string '%v' to find the variable in sequence
	fmt.Printf("my age is %q and my name is %q \n", age, name) //'%q' is used to quote around the variable	when we output them , not going to work with int
	fmt.Printf("age is of type %T \n", age)                    //'%T' is used to get the date type of the variable
	fmt.Printf("your score %f points! \n", 225.55)             //'%f' is used to print float numbers
	fmt.Printf("your score %0.1f points! \n", 225.55)          //'%f' you can limit/round the number points that print with the float numbers

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name) // in here we can save this string format in to a variable
	fmt.Println("the saved string is : ", str)                            // output is the saved string is :  my age is 35 and my name is sandun

}
