package main

import "fmt"

func main() {

	//structs -> blue print which describe type of data(bule print builds that print objects)

	myBill := newBill("mario's bill")

	// fmt.Println(myBill) //this will print bill.go newly created struct {mario's bill map[] 0}

	fmt.Println(myBill.format())
}
