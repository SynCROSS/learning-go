package main

import (
	"fmt"
	"strings"
)

func lenAndUPPER(word string) (length int, UPPERCASE string) {
	defer fmt.Println("I'm done. :)")
	fmt.Println("I'm started!")
	length = len(word)
	UPPERCASE = strings.ToUpper(word)
	return length, UPPERCASE
}
func printElements() {
	fmt.Println(lenAndUPPER("yee"))
}