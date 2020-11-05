package main

import (
	"fmt"
	"learning-go/something"
)

func multiply(num1, num2 int) (int, string) { // * == num1 int , num2 int
	return num1 * num2, "multiplied!"
}
func main() {
	name := "SynCROSS" // * var name string = "Syncross"
	something.SayHi()
	fmt.Println(name) // * To export a function, the function name must begin with a capital letter.
	fmt.Println(multiply(10, 10))
}
