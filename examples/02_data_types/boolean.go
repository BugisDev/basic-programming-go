package main

import "fmt"

func Boolean() {
	var ok bool
	// default is false
	fmt.Println("Default value of bool var is", ok)
	x := 10
	ok = x > 5

	if ok {
		fmt.Println("x is bigger than 5")
	}

	name := "Stranger"
	ok = len(name) > 5
	if ok {
		fmt.Println("your name is too long")
	}
}
