package gobasic

import "fmt"

func SliceRun() []string {
	// create a slice of strings
	slice := []string{"bimo", "dzaki"}
	// add data to slice
	slice = append(slice, "ahmad")

	// create a slice of ints
	ints := []int{1,3,5,6}
	ints = append(ints, 89)

	fmt.Println(slice)
	return slice
}