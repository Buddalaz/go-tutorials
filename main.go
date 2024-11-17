package main

import "fmt"

func main() {

	//structs -> blue print which describe type of data(bule print builds that print objects)

	myBill := newBill("mario's bill")

	myBill.addItem("onion soup", 4.50)
	myBill.addItem("veg pie", 8.95)
	myBill.addItem("toffee pudding", 4.95)
	myBill.addItem("coffee", 3.25)

	myBill.updateBill(10) // this is not going to update the tip in bill bcuz we update the copie that we passed

	// fmt.Println(myBill) //this will print bill.go newly created struct {mario's bill map[] 0}

	fmt.Println(myBill.format())
}
