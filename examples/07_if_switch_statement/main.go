package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	// basic if statement
	if 1 > 2 {
		fmt.Println("1 is greater than 2")
	}

	// basic multiple statement
	if 1 > 2 {
		fmt.Println("1 is greater than 2")
	} else if 2 > 3 {
		fmt.Println("2 is greater than 3")
	}

	// using otherwise or else
	if 1 > 2 {
		fmt.Println("1 is greater than 2")
	} else {
		fmt.Println("1 is less than 2")
	}

	// statement on different types
	if true {
		fmt.Println("Yeah its true")
	}

	var ok bool
	if !ok {
		fmt.Println("too bad, its not ok")
	}

	name := ""
	if name == "" {
		fmt.Println("empty name :(")
	}

	// using high-order function
	var isEmpty = func(s string) bool {
		return s == ""
	}

	if isEmpty(name) {
		fmt.Println("empty name :(")
	}

	// comparing int and uint
	var unsInt uint
	if unsInt == uint(0) {
		fmt.Println("empty number")
	}

	// inline statement, using local variables on the block
	if length := len("ady"); length < 6 {
		// length var only available on this block
		fmt.Printf("your name is too short. length: %d\n", length)
	}

	// using multiple return in if statement
	// convert string `10` into integer using strconv package
	if n, err := strconv.Atoi("10"); err == nil {
		// n and err only available on this block statement
		fmt.Println("n value:", n)
	}

	// unexport result, this should return err
	if _, err := strconv.Atoi("ady"); err != nil {
		fmt.Println("err:", err.Error())
	}

	// we defined our high-order function with multiple returns
	var validate = func(s string) (string, error) {
		if s == "" {
			return s, errors.New("Invalid parameter")
		}
		if len(s) <= 6 {
			return s, errors.New("Invalid Length, should be minimum of 6")
		}
		if len(s) >= 10 {
			return s, errors.New("Invalid Length, should be maximum of 10")
		}
		return s, nil
	}

	// this should error invalid parameter
	if _, err := validate(""); err != nil {
		fmt.Println("err:", err.Error())
	}

	// this should error invalid length
	if _, err := validate("ady"); err != nil {
		fmt.Println("err:", err.Error())
	}

	// this should error invalid length
	if _, err := validate("Ady Rahmat MA"); err != nil {
		fmt.Println("err:", err.Error())
	}

	if s, err := validate("ganteng"); err == nil {
		fmt.Println("s value is validated. s:", s)
	}

	// using switch on expression int
	// no need to use break
	switch len("ady") {
	case 3:
		fmt.Println("length is three")
	default:
		fmt.Println("unknown/undefined length")
	}

	// storing expression value as local variables
	var f interface{}
	f = "ady"
	switch l := f.(type) {
	case int, uint, float64:
		// comma separated means or
		fmt.Printf("f is numeric type. %d\n", l)
	case string:
		fmt.Printf("f is string type. %s\n", l)
	default:
		fmt.Printf("f has unknown/undeclared type. %v\n", l)
	}

	// using switch without expression, means default true
	switch {
	case len("ady") <= 6 || len("ady") >= 10:
		fmt.Println("`ady` has invalid length")
	default:
		fmt.Println("validation succed")
	}
}
