package main

import (
	"fmt"
)

func multiply1(num1, num2 int) (int, string) { // * == num1 int , num2 int
	return num1 * num2, "multiplied!"
}
func function() {
	name := "SynCROSS" // * var name string = "Syncross"
	fmt.Println(name) // * To export a function, the function name must begin with a capital letter.
	fmt.Println(multiply1(10, 10))
}