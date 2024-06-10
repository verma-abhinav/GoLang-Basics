package main

import "fmt"

// Defining a struct type
type Address struct {
	Name    string
	city    string
	Pincode int
}

func structure() {

	// Declaring a variable of a `struct` type
	// All the struct fields are initialized
	// with their zero value
	var a Address
	fmt.Println(a)

	// Declaring and initializing a
	// struct using a struct literal
	a1 := Address{"Akshay", "Dehradun", 3623572}

	fmt.Println("Address1: ", a1)

	// Naming fields while
	// initializing a struct
	a2 := Address{Name: "Anikaa", city: "Ballia",
		Pincode: 277001}

	fmt.Println("Address2: ", a2)

	// Uninitialized fields are set to
	// their corresponding zero-value
	a3 := Address{Name: "Delhi"}
	fmt.Println("Address3: ", a3)

	// Accessing struct fields
	// using the dot operator
	fmt.Println("Address1 Name: ", a1.Name)
	fmt.Println("Address1 city: ", a1.city)

	// Assigning a new value
	// to a struct field
	a1.city = "Bareilly"

	// Displaying the result
	fmt.Println("New Address1: ", a1)

}
