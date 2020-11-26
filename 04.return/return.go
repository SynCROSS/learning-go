package main

import (
	"fmt"
)

func multiply3(num1, num2 int) (int, string) { // * == num1 int , num2 int
	return num1 * num2, "-- multiplied!"
}
func printReturnedElements() {
	result, message := multiply3(10, 10)
	fmt.Println(result, message)
}