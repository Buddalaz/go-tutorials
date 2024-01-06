package main

import "fmt"

var score = 99.5

func main() {
	sayHello("mario") //these function can access from other files

	for _, value := range points {
		fmt.Println(value)
	}

	//but .\main.go:6:2: undefined: sayHello // .\main.go:8:24: undefined: points this will output because of that we have to run those 2 files
	//go run main.go greeting.go

	// var score = 99.5 can't access this inside the main function it has to be outside the function

	showScore()

}
