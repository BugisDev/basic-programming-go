package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

func main() {
	// in go, there are no exceptions
	// in go, error handled using return values

	// try to use strconv package from go
	// we are going to convert x into integer
	x := "name"
	n, err := strconv.Atoi(x)
	if err != nil {
		log.Printf("x cannot be converted to an integer. err: %s\n", err.Error())
	} else {
		log.Printf("n: %d is integer value of x\n", n)
	}

	// err is an interface value from "error" interface
	// and because its an interface, we can define our own error

	// using package errors to define error
	err = check(x)
	if err != nil {
		log.Printf("got err: %s\n", err.Error())
	}

	// using defined variables, it is very usefull
	// like catching exception on other language
	err = anotherCheck(x)
	if err == ErrInvalidLength {
		log.Println("X is less than 5 char")
	}

	// using type struct, which we can got more information about the error
	// we can do type assertion here
	err = checkString(x)
	terr, ok := err.(MyError)
	if ok {
		log.Println(terr.Error())
		log.Printf("%d: %s\n", terr.Code, terr.Message)
	}
}

func check(x string) error {
	if len(x) < 5 {
		return errors.New("Invalid Length")
	}
	return nil
}

var ErrInvalidLength = errors.New("Invalid Length")

func anotherCheck(x string) error {
	if len(x) < 5 {
		return ErrInvalidLength
	}
	return nil
}

type MyError struct {
	Code    int
	Message string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

func checkString(x string) error {
	if len(x) < 5 {
		return MyError{Code: 409, Message: "Invalid Length"}
	}
	return nil
}
