package main

import (
	"fmt"
)

func addSomething(numbers ...int) int {
	total := 0
	for _, number := range numbers { // * you can define variable in for, switch, etc
		total += number
	}
	return total
}

func printElements() {
	fmt.Println(addSomething(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
