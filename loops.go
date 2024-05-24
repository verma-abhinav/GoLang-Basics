package main

import "fmt"

func loops() {

	// WHILE LOOP
	x := 0
	for x < 5 {
		fmt.Println("Value of x is:-", x)
		x++
	}

	fmt.Println("------------------------------")

	// FOR LOOP
	for i := 0; i < 5; i++ {
		fmt.Println("Value of i is:-", i)

	}

	fmt.Println("------------------------------")

	// FOR LOOP with SLICES
	names := []string{"Mario", "Goku", "Luigi", "Trent", "Vegeta"}
	for i := 0; i < len(names); i++ {
		fmt.Printf("Name of %vth person is %v. \n", i, names[i])
	}

	fmt.Println("------------------------------")

	// FOR LOOP with SLICES using range
	for index, value := range names {
		fmt.Printf("Hello %vth person %v. \n", index, value)
	}

	fmt.Println("------------------------------")

	// FOR LOOP with SLICES using range with either index or value
	for _, value := range names {
		fmt.Printf("Val of person is %v. \n", value)
	}
}
