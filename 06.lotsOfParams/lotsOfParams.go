package main

import (
	"fmt"
)

func repeat(words ...string) {
	fmt.Println(words)
}
func lotsOfParams() {
	repeat("y", "e", "e")
}