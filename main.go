package main

import "fmt"

func updateName(x string) {
	x = "wedge"
}

func updateNameThroughPointer(x *string) { //accepting a pointer what ever passing to the paramater value store in the memory location
	*x = "wedge" //dereference the passing pointer into the variable
}

func main() {

	//pointers -> a pointer is just a pointer to a memory location

	name := "tifa"

	updateName(name)

	fmt.Println("memory address of name is: ", &name) //-> create a pointer to the memory location of the name vaiable and the output is like
	//memory address of name is: 0xc000050250

	//store pointers into a variable

	m := &name //create a variable called m and ampersion the name into it(storing the memory location pointer in its own memory blocks)

	fmt.Println("memory address: ", m) //this will print the same memory address of the name memory address: 0xc000050250

	fmt.Println(name)

	//to access the value of the pointer we use astric before the pointer variable

	fmt.Println("value at memory address: ", *m) //value at memory address: tifa

	fmt.Println(name)

	//change the value of the pointer by passing into a function

	updateNameThroughPointer(m)

	fmt.Println(name) //will print the updated value as wedge

}

/*
|--name---|----m----|
|  0x001  |  0x002  |
|---------|---------|
| "tifa"  |  p0x001 |
|---------|---------|
*/
