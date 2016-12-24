package main

import "fmt"

// exported var
var GlobalVar int

func main() {
	var x int
	x = 1
	var y uint
	y = uint(x)
	var (
		a, b int
		name string
	)
	a, b = x, int(y)
	name = "Stranger"
	hi := "Hellooo " + name

	zy := 10
	fmt.Printf("Type of zy: %T\n", zy)

	printInterface(a)
	printInterface(b)
	printInterface(hi)

	printGlobal()
	updateGlobal()
	printGlobal()
}

func printNumeric(x int) {
	fmt.Println(x)
}

func printInterface(x interface{}) {
	fmt.Println("x :", x)
}

func printGlobal() {
	fmt.Println("GlobalVar", GlobalVar)
}

func updateGlobal() {
	GlobalVar = 10
}
