package main

import (
	"fmt"
	"strings"
)

// return multiple values from a function
func getInitials(value string) (string, string) { //define the sequence/order to whats returns
	//convert the string into capitals
	s := strings.ToUpper(value)

	//split the two words into a slice
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"

}

func main() {

	//calling the function
	fName1, Lname1 := getInitials("sandun sampath")

	fmt.Println(fName1, Lname1) //S S

	fName2, Lname2 := getInitials("cads ddsads")

	fmt.Println(fName2, Lname2) //C D

	fName3, Lname3 := getInitials("sdsads")

	fmt.Println(fName3, Lname3) //S _

}
