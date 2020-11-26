package main

import (
	"fmt"
)

func multiply2(num1, num2 int) (int, string) { // * == num1 int , num2 int
	return num1 * num2, "-- multiplied!"
}
func ignoredValue() {
	result, _ := multiply2(10, 10)
	fmt.Println(result)
}