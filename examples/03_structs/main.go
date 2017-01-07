package main

import "fmt"

type Person struct {
	ID   int
	Name string
}

func main() {
	// define new person
	var ganteng Person

	// assign
	ganteng = Person{ID: 1, Name: "Si Ganteng"}

	// without key
	ganteng = Person{1, "Si Ganteng"}

	// without anything
	ganteng = Person{}
	ganteng.ID = 1
	ganteng.Name = "Si Ganteng"

	// create new person
	cantik := Person{2, "Si Cantik"}

	fmt.Printf("%d. %s\n", ganteng.ID, ganteng.Name)
	fmt.Printf("%d. %s\n", cantik.ID, cantik.Name)

	// inline struct
	var anotherPerson struct {
		ID   int
		Name string
		Job  string
	}

	anotherPerson.ID = 3
	anotherPerson.Name = "Entah siapa"
	anotherPerson.Job = "None"

	fmt.Printf("%d. %s\n", anotherPerson.ID, anotherPerson.Name)

	newPerson := NewPerson{Person{ID: 4, Name: "New Person"}}
	fmt.Printf("%d. %s\n", newPerson.ID, newPerson.Name)

	// using defined method
	fmt.Printf("%d. %s\n", newPerson.ID, newPerson.WhoAmI())

	// change my name
	newPerson.ChangeMyName("Call me si Ganteng")
	fmt.Printf("%d. %s\n", newPerson.ID, newPerson.WhoAmI())

	fmt.Println()
	fmt.Printf("%s is %d long\n", ganteng.Name, HowLongMyName(ganteng))
	fmt.Printf("%s is %d long\n", newPerson.Name, HowLongMyName(newPerson.Person))

	fmt.Println()
	var people People
	people.AddNewPerson(ganteng)
	fmt.Printf("We got %d people\n", people.Len())
	people.AddNewPerson(cantik)
	fmt.Printf("We got %d people\n", people.Len())
	people.AddNewPerson(newPerson.Person)
	fmt.Printf("We got %d people\n", people.Len())

	fmt.Println()
	var community = Community{Name: "Kumpulan Orang Ganteng dan Cantik", People: people}
	fmt.Printf("%s have %d people\n", community.Name, community.TotalPerson())
	community.AddPerson(Person{5, "Sedikit Ganteng"})
	fmt.Printf("%s have %d people\n", community.Name, community.TotalPerson())
}

// NewPerson will have all properties and methods on Person struct
type NewPerson struct {
	Person
}

// WhoAmI tell me who i am (value)
func (np NewPerson) WhoAmI() string {
	return np.Name
}

// ChangeMyName change Name value (pointers)
func (np *NewPerson) ChangeMyName(newName string) {
	np.Name = newName
}

// HowLongMyName count the length of person.name
func HowLongMyName(person Person) int {
	return len(person.Name)
}

// People is a bunch of person (slice of Person : similar to Array)
type People []Person

func (p *People) AddNewPerson(person Person) {
	*p = append(*p, person)
}

func (p People) Len() int {
	return len(p)
}

// Community contain name and list of person from People structure
type Community struct {
	Name string
	People People
}

// AddPerson compositing AddNewPerson from People
func (c *Community) AddPerson(person Person) {
	c.People.AddNewPerson(person)
}

// TotalPerson compositing Len function from People
func (c Community) TotalPerson() int{
	return c.People.Len()
}
