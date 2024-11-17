package main

import "fmt"

//create a struct
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

//make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip:   0,
	}
	return b
}

//add methods as receiver fuction to the as associate with bill
//format the bill -> where we receive bill object into this fuction can access that within the fuction
func (b bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	//list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v) //in -25 means it will add 25 character space from its left character length. if its plus 25 it will add from right
		total += v
	}

	//total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)

	return fs
}
