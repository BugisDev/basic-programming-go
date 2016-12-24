package main

import "fmt"

func Interface() {
	var x interface{}
	x = 10
	x, ok := x.(int)
	if ok {
		fmt.Println("Yeah, x is an integer")
	}

	var y interface{}
	y = "Stranger"
	y, ok = y.(int)
	if !ok {
		fmt.Println(":(, so sad. y is not an integer")
	}
}
