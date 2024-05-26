package main

import "fmt"

func maps() {
	//Creating and storing data in maps
	menu := map[string]float64{
		"Soup":      4.99,
		"Pie":       5.99,
		"Sandwitch": 7.99,
	}
	fmt.Println(menu)
	fmt.Println(menu["Soup"])
	//Iterating through a map
	for key, val := range menu {
		fmt.Printf("The cost of %v is %v \n", key, val)
	}
	menu["Pie"] = 11.99
	fmt.Println(menu)
	menu["Noodles"] = 15.99
	fmt.Println(menu)
}
