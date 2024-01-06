package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	//go different libries
	//strings packeage

	greeting := "hello there friends!"
	fmt.Println(strings.Contains(greeting, "hello")) //in here strings is a libery that import from go this will check the given search value is
	// contains in the greeting word, if exist return true else false
	fmt.Println(strings.Contains(greeting, "hello!")) //this retruns false

	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi")) //output -> hi there friends!
	//in here in the string it replace the all value hello with hi. this dosen't override the greeting
	//string this will create a new string with the replace once

	fmt.Println(strings.ToUpper(greeting)) //this will print all the words with capitals output-> HELLO THERE FRIENDS!

	fmt.Println(strings.Index(greeting, "ll")) //find the index of the provided word output-> 2

	fmt.Println(strings.Split(greeting, " ")) //split the words in to an array output-> [hello there friends!]

	fmt.Println("original string value = ", greeting) //output -> original string value =  hello there friends!

	//sort package

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	//sort slice of integers
	sort.Ints(ages)   //this method altered the original slice
	fmt.Println(ages) //output -> [20 25 30 35 45 50 60 75]

	index := sort.SearchInts(ages, 30) //search slice of integers, return the position of 30
	fmt.Println(index)                 //output will be get from the previous result sorted output -> 2

	names := []string{"animal", "bird", "fish", "reptailes"}
	sort.Strings(names) //sort this slice in to alphabetical order
	fmt.Println(names)  //output -> [animal bird fish reptailes]

	fmt.Println(sort.SearchStrings(names, "fish")) //search string and return the index of the sorted slice, output -> 2

}
