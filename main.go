package main

import (
	"fmt"
	"strings"
)

func lenAndUPPER(word string) (length int, UPPERCASE string) {
  	defer fmt.Println("I'm done. :)")
	length = len(word)
	UPPERCASE = strings.ToUpper(word)
	return length, UPPERCASE
}
func main() {
	fmt.Println(lenAndUPPER("yee"))
}
