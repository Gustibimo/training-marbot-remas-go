package gobasic

import "fmt"

// struct can contains  multiple values of different data types - composite data type
type Person struct {
	Name string
	Age int
	Gender string
}

func StructRun()  {

	// instantiate Person
	p := Person{
		Name: "Umar",
		Age: 40,
		Gender: "Male",
	}

	// alternative way to instatiate
	p2 := new(Person)
	p2.Name = "Ismael"
	p2.Age = 30
	p2.Gender = "Male"

	fmt.Println(p)
}