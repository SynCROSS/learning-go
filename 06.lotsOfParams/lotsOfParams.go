package main

import (
	"fmt"
)

func repeat(words ...string) { // * ... can get same type values
	fmt.Println(words)
}
func lotsOfParams() {
	repeat("y", "e", "e")
}
