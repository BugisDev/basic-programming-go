package main

import "fmt"

func main() {
	// define slice of integers
	var numbers []int

	numbers = []int{1, 2, 3, 4}
	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Len: %d, Capacity: %d\n", len(numbers), cap(numbers))

	// adding more numbers
	numbers = append(numbers, 2)
	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Len: %d, Capacity: %d\n", len(numbers), cap(numbers))
	// or we can use
	numbers = append(numbers, 6, 7)
	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Len: %d, Capacity: %d\n", len(numbers), cap(numbers))

	// accessing value
	number := numbers[6]
	fmt.Printf("Number of index 10 = %d\n", number)

	// accessing unknown index
	if len(numbers) > 500 {
		number = numbers[499]
		fmt.Printf("Number of index 499 = %d\n", number)
	}

	// deleting data from slice
	deleteIndex := 2
	numbers = append(numbers[:deleteIndex], numbers[deleteIndex+1:]...)
	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Len: %d, Capacity: %d\n", len(numbers), cap(numbers))

	var names []string

	names = append(names, "Ady Rahmat MA", "Nur Yaumi")
	names = append(names, "Eka Putra", "Chaidir Yahya")
	names = append(names, "Herianto", "Asriadi N", "Risyad Rais")
	fmt.Printf("List of names %v\n", names)

	// slicing a slice
	listNames := names[0:2]
	fmt.Printf("List of names %v\n", listNames)

	// slicing from behind
	listNames = names[len(names)-2:]
	fmt.Printf("List of names %v\n", listNames)

	listNames = names[0:2:4]
	fmt.Printf("List of names %v\n", listNames)

	// what if we had a slice of string and integer, or any other type?
	var data []interface{}
	data = append(data, 1)
	data = append(data, "Ady")
	data = append(data, nil)
	data = append(data, func() {
		fmt.Println("hello")
	})
	data = append(data, map[string]string{"name": "ady"})
	fmt.Printf("Data %v\n", data)

	// nested slice
	var communities [][]string
	names = []string{"ady", "yaumi"}
	communities = append(communities, names)
	fmt.Printf("Communities %v\n", communities)
	names = []string{"eka", "herianto"}
	communities = append(communities, names)
	fmt.Printf("Communities %v\n", communities)
}
