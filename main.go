package main

import "fmt"

func main() {

	//maps
	//maps allow us to store key value pairs whether keys can be different types and the vlues different types as well, 
	//but in a single mapp all of the keys must be same type and all of the values must be the same type as well
	fmt.Println("mapp excerisce.....")

	//create a map
	menu := map[string]float64{
		"soup":          4.99,
		"pie":           7.99,
		"salad":         6.99,
		"tofee pudding": 3.55,
	}

	//print the map
	fmt.Println(menu) //map[pie:7.99 salad:6.99 soup:4.99 tofee pudding:3.55]

	//print a key in the map
	fmt.Print(menu["pie"]) //7.99

	//looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	//ints as key type
	phoneBooks := map[int]string{
		1234: "sandun",
		5678: "sampath",
		9123: "dewage",
	}

	fmt.Println(phoneBooks) //map[1234:sandun 5678:sampath 9123:dewage]

	fmt.Println(phoneBooks[1234]) //sandun

	//update item inside a map
	phoneBooks[9123] = "ginimala" //can't assign a int cuz key can't change and should be assign a define type which is string
	fmt.Println(phoneBooks) //map[1234:sandun 5678:sampath 9123:ginimala]

	phoneBooks[1234] = "eppawela"
	fmt.Println(phoneBooks) //map[1234:eppawela 5678:sampath 9123:ginimala]

}
