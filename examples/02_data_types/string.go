package main

import (
	"fmt"
	"strconv"
	"strings"
)

func String() {
	var name string
	// default is "" or empty (not null/nil)
	fmt.Println("Default value of string is", name)
	name = "Stranger"

	fmt.Println("Calling", name)
	fmt.Println("Len of name", len(name))
	// slice 0,3
	fmt.Println("Slicing name", name[:3])
	// slice -3
	fmt.Println("Slicing name", name[3:])
	// slice 3,5
	fmt.Println("Slicing name", name[3:5])

	fmt.Println("lowercasing name", strings.ToLower(name))
	fmt.Println("uppercasing name", strings.ToUpper(name))

	var x int

	n := "10"
	// convert string to integer
	x, _ = strconv.Atoi(n)
	fmt.Println("Now x:", x)
}
