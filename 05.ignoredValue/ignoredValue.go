package main

import (
	"fmt"
)

func multiply2(num1, num2 int) (int, string) {
	return num1 * num2, "-- multiplied!"
}
func ignoredValue() {
	result, _ := multiply2(10, 10) // * _ can ignore value
	fmt.Println(result)
}
