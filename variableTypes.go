package main

import "fmt"

func updateName(name string) string {
	name = "Vegeta"
	return name
}
func updateMenu(menu map[string]float64) {
	menu["Pie"] = 15.99
}
func variableTypes() {
	// non-pointer types -> String, Int, Float, Boolean, Arrays, Structs
	name := "Goku"
	name = updateName(name)
	fmt.Println(name)
	// pointer wrapper types -> map, slice, function
	menu := map[string]float64{
		"soup":      4.99,
		"Pie":       5.99,
		"Sandwitch": 7.99,
	}
	updateMenu(menu)
	fmt.Println(menu)
}
