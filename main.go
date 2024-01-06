package main

import (
	"fmt"
	"math"
)

// define a function in outside the main function
func sayGreeting(value string) {
	fmt.Printf("Good morning %v \n", value)
}

// single arguments
func sayBye(value string) {
	fmt.Printf("Goodbye %v \n", value)
}

// multiple arguments
func cycleNames(s []string, f func(string)) { //function can be also pass as an arguments in this second argument accepted function is string
	for _, v := range s {
		f(v)
	}
}

// return value from a function
func circleArea(r float64) float64 { //in here we specify the return value after the parenthesis
	return math.Pi * r * r
}

func main() {

	//calling a function
	// sayGreeting("mario") //output -> Good morning mario
	// sayGreeting("luigi") //output -> Good morning luigi
	// sayBye("mario")      //output -> Goodbye mario

	//calling multiple arguments
	cycleNames([]string{"simon", "ksi", "vik123"}, sayGreeting) //we only pass the reference in the calling function
	cycleNames([]string{"simon", "ksi", "vik123"}, sayBye)      //we only pass the reference in the calling function
	//output is
	// Good morning simon
	// Good morning ksi
	// Good morning vik123
	// Goodbye simon
	// Goodbye ksi
	// Goodbye vik123

	//return value from a function
	a1 := circleArea(2.34)
	a2 := circleArea(10)

	fmt.Println("Area 1 is", a1) //output -> Area 1 is 17.20210473399627
	fmt.Println("Area 2 is", a2) //output -> Area 2 is 314.1592653589793

	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f", a1, a2) //output -> circle 1 is 17.202 and circle 2 is 314.159
}
