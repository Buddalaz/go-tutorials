package main

import "fmt"

func updateName1(x string) {
	x = "wedge"
}

func updateName2(x string) string {
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main() {

	//pass-by-value
	//Go makes "copies" of values when passed into functions

	//group A types(non-pointer values) -> strings, ints, bools, floats, arrays, structs
	name := "tifa"

	updateName1(name)

	fmt.Println(name) //print tifa

	name = updateName2(name) //in this method will return the updated copie value

	fmt.Println(name) //print wedge

	//group B types(pointer wrapper Values) -> slices, maps, functions

	menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}

	updateMenu(menu) //when passing the map(pointer wrapper Values) this will update/change the same memory block

	fmt.Println(menu) //print map[coffee:2.99 ice cream:3.99 pie:5.95]

}
