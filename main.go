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
	word, e2 := dictionary.SearchWord("name")
	if e2 != nil {
		fmt.Println(e)
	}
	dictionary.DeleteWord("name")
	word, e2 = dictionary.SearchWord("name")
	if e2 != nil {
		fmt.Println(e2)
	}
	fmt.Println(word)
}
