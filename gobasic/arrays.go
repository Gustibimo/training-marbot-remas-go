package gobasic

import "fmt"

func ArrayRun() {

	// Declare array of integer with specific size
	var array [5]int
	array[0] = 10

	// Declare array with data
	arrayWithData := [4]string{"ahmad", "Dzaki", "Dani"}

	// Add data to defined array
	arrayWithData[3] = "Bimo"

	fmt.Println(arrayWithData)

}
