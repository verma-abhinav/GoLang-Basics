package main

import "fmt"

func sayGreeting(name string) {
	fmt.Printf("Greetings %v! How are you today? \n", name)

}
func sayBye(name string) {
	fmt.Printf("Bye %v! See you again. \n", name)

}
func cycleNames(name []string, callmyfun func(string)) {
	for _, val := range name {
		callmyfun(val)
	}
}
func calArea(radius float64) float64 {
	return 3.14 * radius * radius
}
func functionsInGo() {
	// sayGreeting("John")
	// sayBye("John")
	cycleNames([]string{"Goku", "Vegeta"}, sayGreeting)
	cycleNames([]string{"Goku", "Vegeta"}, sayBye)
	area := calArea(7)
	fmt.Println("Area of circle with radius 7 is ", area)
}
