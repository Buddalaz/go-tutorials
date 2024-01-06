package main

import "fmt"

//in this greeing.go still in this package main we can access these functions, variables including thr main.go

var points = []int{20, 90, 100, 45, 70}

func sayHello(value string) {
	fmt.Println("hello", value)
}

func showScore() {
	fmt.Println("you score this many points: ", score)
}
