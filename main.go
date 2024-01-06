package main

import "fmt"

func main() {

	//create an array
	var ages1 [3]int = [3]int{20, 25, 30} //this array has fixed leangth
	var ages2 = [3]int{10, 20, 30}        //this array is same as this

	//print array elements and length
	fmt.Println(ages1, len(ages1))
	fmt.Println(ages2, len(ages2)) // their is a function for check leangth of an array called len()

	//string arrays
	names := [4]string{"adb", "asc", "sdas", "sasrfe"}
	names[1] = "sadun" //can change element values
	fmt.Println(names, len(names))

	//slices (use arrays under the hood)
	//we can change leangth, can add items to it or take items awey with it
	var scores = []int{100, 50, 60} //not define the leangth
	scores[2] = 25

	//append item to this slices
	scores = append(scores, 85) //what happends here was, this dosent cahnge the scores instead of this return new slices with the append score just like override the scores

	fmt.Println(scores, len(scores))

	// slices ranges -> range is a way to get range of element from existing arrays wishe and store them to slice
	rangeOne := names[1:3]
	fmt.Println(rangeOne) //output ->[sadun sdas]

	rangeTwo := names[2:] //this will get index 2 to end of the items
	fmt.Println(rangeTwo) //output ->[sdas sasrfe]

	rangeThree := names[:3] //this will get from start to index 3 not the 3
	fmt.Println(rangeThree) //output -> [adb sadun sdas]

	//we can also append items to ranges end of the day they also slices
	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne) //output -> [sadun sdas koopa]

}
