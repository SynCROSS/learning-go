package main

import (
	"fmt"
)

func main() {
	a := 10
	b := &a // * Modify b to modify a
	*b *= 10
	fmt.Println(a)
}
