package main

import "fmt"

func main() {

	//boolean

	age := 45

	fmt.Println(age <= 50) //age is grater than or equal 50, output-> true
	fmt.Println(age >= 50) //age is less than or equal 50, output-> false
	fmt.Println(age == 45) //age is equal 45, output-> true
	fmt.Println(age != 50) //age is not 50, output-> true

	//conditionals
	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is not less than 45")
	}

	names := []string{"aniaml", "birds", "fish", "reptails"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}
		fmt.Printf("the value at pos %v is %v \n", index, value)
	}
	//output is
	// 	the value at pos 0 is aniaml
	// continuing at pos 1
	// the value at pos 2 is fish
	// the value at pos 3 is reptails

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println("breaking at pos", index)
			break
		}
		fmt.Printf("the value at pos %v is %v \n", index, value)
	}
	//output is
	// 	the value at pos 0 is aniaml
	// continuing at pos 1
	// the value at pos 2 is fish
	// breaking at pos 3

}
