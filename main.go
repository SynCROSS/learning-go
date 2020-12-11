package main

import (
	"fmt"
	"main/dict"
)

func main() {
	dictionary := dict.Dictionary{}
	e := dictionary.AddWord("name", "SynCROSS")
	if e != nil {
		fmt.Println(e)
	}
	definition, _ := dictionary.Search("name")
	fmt.Println(definition)
	e2 := dictionary.AddWord("name", "SynCROSS")
	if e2 != nil {
		fmt.Println(e2) // * Error
	}
}
