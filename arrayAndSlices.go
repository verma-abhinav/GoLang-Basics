package main

import (
	"fmt"
)

func arrayAndSlices() {
	//Arrays
	var ages1 [3]int = [3]int{25, 30, 35}
	var ages2 = [3]int{45, 50, 55}
	ages3 := [3]int{70, 75, 80}
	ages4 := [...]int{80, 85, 90, 100, 110}
	var ages5 = [...]int{20, 25, 30, 35, 40, 45, 50, 55}

	fmt.Println(ages1, "\n", ages2, len(ages2), "\n", ages3, "\n", ages4, len(ages4), "\n", ages5)

	//Slices
	var mySlice1 = []int{20, 30, 40}
	mySlice2 := []int{50, 60, 70}
	fmt.Println(mySlice1, len(mySlice1))
	fmt.Println(mySlice2, len(mySlice2), cap(mySlice2))

	// Slices from Arrays
	arr1 := [6]int{30, 40, 50, 60, 70, 80}
	mySlice3 := arr1[2:4]
	fmt.Println(mySlice3, len(mySlice3), cap(mySlice3))

	// with make() function
	mySlice4 := make([]int, 5, 10)
	fmt.Println(mySlice4, len(mySlice4), cap(mySlice4))

}
