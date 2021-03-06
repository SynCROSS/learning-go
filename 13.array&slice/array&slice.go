package main

import (
	"fmt"
)

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // * This is the array in Go
	slice := []string{"s", "l", "i", "c", "e"}      // * This is the Slice in Go
	slice = append(slice, "🔪")                      // * NOT slice.append() Because append returns new array

	fmt.Println(array)
	fmt.Println(slice)
}
