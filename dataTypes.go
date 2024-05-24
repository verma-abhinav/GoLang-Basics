package main

import "fmt"

func dataTypes() {
	//Strings

	var NameOne string = "Max"
	var NameTwo = "Page"
	NameThree := "Hope"
	var NameFour string
	fmt.Println(NameOne, NameTwo, NameThree, NameFour)
	NameFour = "Hydra"
	fmt.Println(NameOne, NameTwo, NameThree, NameFour)

	// Integers
	var NumOne int = 25
	var NumTwo = 30
	NumThree := 40
	var NumFour int
	fmt.Println(NumOne, NumTwo, NumThree, NumFour)
	NumFour = 50
	fmt.Println(NumOne, NumTwo, NumThree, NumFour)

}
