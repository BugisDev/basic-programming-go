package main

import "fmt"

func main() {
	// basic loop
	for i := 0; i < 10; i++ {
		// i only available in this block
		fmt.Printf("%d", i)
	}
	fmt.Println()

	// loop on types expression
	var pleaseLoop bool = true
	for pleaseLoop {
		fmt.Println("loop yeah")
		pleaseLoop = !pleaseLoop
	}

	// theres no while, just for
	// for is while, use break to break
	var i int
	for {
		// we use i as incrementor
		if i > 9 {
			// break the loop
			break
		}
		fmt.Printf("%d", i)
		i++
	}
	fmt.Println()

	// using continue to skipping a loop
	for i := 0; i < 10; i++ {
		if i == 4 || i == 6 {
			continue
		}
		fmt.Printf("%d", i)
	}
	fmt.Println()

	// loop a slice
	var numbers []int
	numbers = []int{1, 2, 3, 4}
	for index, number := range numbers {
		// index and number are local variables, only available on this block statement
		// index are always a number
		fmt.Printf("Index: %d. Value: %d\n", index, number)
	}

	// looping without the index
	for _, number := range numbers {
		fmt.Print(number)
	}
	fmt.Println()

	var names []Person
	names = append(names, Person{"Ady"}, Person{"Nur Yaumi"})
	for index, p := range names {
		// index are number, p are person type
		fmt.Printf("%d. %s\n", index+1, p.Speak())
	}

	var peoples []CanSpeak
	peoples = append(peoples, names[0])
	peoples = append(peoples, AnotherPerson{"Eka", "College Student"})
	for index, p := range peoples {
		// index are integer, p are interface that implement CanSpeak interface
		fmt.Printf("%d. %s\n", index+1, p.Speak())
	}

	// loop on a map
	var myMap map[string]interface{}
	myMap = map[string]interface{}{
		"book":   "Buku",
		"satu":   1,
		"person": names[0],
	}
	for key, value := range myMap {
		// just print the key and the value
		fmt.Printf("%s: %v\n", key, value)
	}

	// switching by type
	for key, value := range myMap {
		// we need local var to store the result of type assertion
		switch v := value.(type) {
		case int:
			fmt.Printf("%s: %d\n", key, v)
		case string:
			fmt.Printf("%s: %s\n", key, v)
		case Person:
			fmt.Printf("%s: %s\n", key, v.Speak())
		default:
			// unknown types, undeclared on case
			fmt.Printf("%s: %v\n", key, value)
		}
	}
}

type Person struct {
	Name string
}

func (p Person) Speak() string {
	return fmt.Sprintf("Hi, my name is %s", p.Name)
}

type AnotherPerson struct {
	Name string
	Job  string
}

func (p AnotherPerson) Speak() string {
	return fmt.Sprintf("Hi, my name is %s. And my i'm a %s", p.Name, p.Job)
}

type CanSpeak interface {
	Speak() string
}
