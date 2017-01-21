package main

import (
	"errors"
	"fmt"
)

func main() {
	// calling function
	sayHi("ganteng")

	// calling multiple param function
	concatName("ady", "ganteng")
	concatNameAgain("ady", "ganteng")

	var name string
	// calling function who return value
	name = concatNameReturn("ady", "ganteng")
	sayHi(name)
	name = concatNameDefinedReturn("ady", "ganteng")
	name = concatNameDefinedReturnAgain("ady", "ganteng")

	var x, y int
	x, y = swapMultipleReturn(10, 15)
	fmt.Printf("x: %d, y: %d\n", x, y)
	x, y = swapWithDefined(y, x)
	fmt.Printf("x: %d, y: %d\n", x, y)
	x, y = swapWithDefinedAgain(y, x)
	fmt.Printf("x: %d, y: %d\n", x, y)

	person, err := returningStruct(name)
	if err != nil {
		fmt.Println("error", err.Error())
	}
	fmt.Printf("Name %s\n", person.Name)

	_, err = returningStruct(name)
	if err != nil {
		fmt.Println("error", err.Error())
	}

	person, _ = returningStruct(name)
	fmt.Printf("Name %s\n", person.Name)

	_, err = pointerParam(nil)
	if err != nil {
		fmt.Println("error", err.Error())
	}

	_, err = pointerParamAndDereference(nil)
	if err != nil {
		fmt.Println("error", err.Error())
	}

	person, _ = pointerParam(&name)
	fmt.Printf("Name %s\n", person.Name)

	interfaceParam(name)
	interfaceParam(x)
	interfaceParam(uint(y))
	interfaceParam(nil)

	multipleParamAsSlice("ady", "ganteng")

	funcParam(func() string {
		return "ady ganteng"
	})
}

func sayHi(name string) {
	fmt.Println("Hi", name)
}

func concatName(firstName string, lastName string) {
	sayHi(firstName + " " + lastName)
}

func concatNameAgain(firstName, lastName string) {
	sayHi(firstName + " " + lastName)
}

func concatNameReturn(firstName, lastName string) string {
	return firstName + " " + lastName
}

func concatNameDefinedReturn(firstName, lastName string) (fullName string) {
	fullName = firstName + " " + lastName
	return fullName
}

func concatNameDefinedReturnAgain(firstName, lastName string) (fullName string) {
	fullName = firstName + " " + lastName
	return
}

func swapMultipleReturn(x, y int) (int, int) {
	return y, x
}

func swapWithDefined(x, y int) (z int, a int) {
	z = y
	a = x
	return
}

func swapWithDefinedAgain(x, y int) (z, a int) {
	z = y
	a = x
	return
}

type Person struct {
	Name string
}

func returningStruct(name string) (person Person, err error) {
	if name == "" {
		return person, errors.New("Empty Name")
	}

	person.Name = name
	return person, nil
}

func pointerParam(name *string) (person Person, err error) {
	if name == nil {
		return person, errors.New("Invalid Name")
	}
	if *name == "" {
		return person, errors.New("Empty Name")
	}

	person.Name = *name
	return person, nil
}

func pointerParamAndDereference(name *string) (person Person, err error) {
	if name == nil {
		return person, errors.New("Invalid Name")
	}
	return returningStruct(*name)
}

func interfaceParam(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Printf("value is string. value: %s\n", v)
	case int, uint:
		fmt.Printf("value is numeric. value: %d\n", v)
	case nil:
		fmt.Println("value is nil")
	}
}

func multipleParamAsSlice(names ...string) {
	for _, name := range names {
		sayHi(name)
	}
}

func funcParam(f func() string) {
	sayHi(f())
}
