package main

import (
	"fmt"
)

func multiply3(num1, num2 int) (int, string) {
	return num1 * num2, "-- multiplied!" // * return values
}
func printReturnedElements() {
	result, message := multiply3(10, 10)
	fmt.Println(result, message)
}
