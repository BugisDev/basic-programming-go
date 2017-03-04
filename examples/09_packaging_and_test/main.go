package main

import (
	"fmt"

	"github.com/bugisdev/basic-programming-go/examples/09_packaging_and_test/slicer"
)

func main() {
	var data = []int{1, 2, 3, 4}
	fmt.Println(data)
	var anotherData = slicer.Int{1, 2, 3, 4}
	fmt.Println(anotherData)

	// checking value (2) exist on data
	for _, value := range data {
		if value == 2 {
			// fire up, value exist
			fmt.Println("2 exist on data")
		}
	}

	// on anotherData
	if anotherData.IsExist(2) {
		// fire up, value exist
		fmt.Println("2 exist on anotherdata")
	}

	// adding value
	data = append(data, 10)
	anotherData.Add(10)
}
