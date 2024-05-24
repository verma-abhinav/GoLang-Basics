package main

import (
	"fmt"
	"sort"
	"strings"
)

func stdLibrary() {
	// string function:-original value remains same,new string is created for output
	greetings := "Hello, how are you doing ?"
	fmt.Println(strings.Contains(greetings, "Hello"))
	fmt.Println(strings.ToUpper(greetings))
	fmt.Println(strings.ReplaceAll(greetings, "Hello", "Hi"))
	fmt.Println(strings.Index(greetings, "ll"))
	fmt.Println(strings.Split(greetings, " "))

	//Sort function:-original value is changed
	ages := []int{145, 555, 265, 175, 485, 395, 105}
	sort.Ints(ages)
	fmt.Println(ages)
	index := sort.SearchInts(ages, 265)
	fmt.Println(index)
}
