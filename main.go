package main

import (
	"fmt"
	"learning-go/something"
)

func multiply(num1, num2 int) (int, string) { // * == num1 int , num2 int
	return num1 * num2, "-- multiplied!"
}
func main() {
	result, _ := multiply(10, 10) // * var name string = "Syncross"
	something.SayHi()
	fmt.Println(result)
}
