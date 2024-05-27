package main

import "fmt"

func pointers() {

	// taking a normal variable
	var x int = 5748

	// declaration of pointer
	var p *int

	// initialization of pointer
	p = &x

	// displaying the result
	fmt.Println("Value stored in x = ", x)
	fmt.Println("Address of x = ", &x)
	fmt.Println("Value stored in variable p = ", p)
	fmt.Println("--------------------------------")

	// taking a pointer
	var s *int

	// displaying the result
	fmt.Println("s = ", s)
	fmt.Println("--------------------------------")
	// using := operator to declare
	// and initialize the variable
	y := 458

	// taking a pointer variable using
	// := by assigning it with the
	// address of variable y
	q := &y

	fmt.Println("Value stored in y = ", y)
	fmt.Println("Address of y = ", &y)
	fmt.Println("Value stored in pointer variable q = ", q)
	fmt.Println("Value stored in y(*q) = ", *q)
	*q = 600
	fmt.Println("Value stored in y(*q) after Changing = ", y)
}
