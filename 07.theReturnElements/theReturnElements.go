package main

import (
	"fmt"
	"strings"
)

func lenAndUPPER(word string) (length int, UPPERCASE string) {
	length = len(word)
	UPPERCASE = strings.ToUpper(word)
	return length, UPPERCASE // * you can set return value
}
func printElements() {
	fmt.Println(lenAndUPPER("yee"))
}