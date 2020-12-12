package main

import (
	"fmt"
	"log"
	"main/dict"
)

func main() {
	dictionary := dict.Dictionary{}
	e := dictionary.AddWord("name", "SynCROSS")
	if e != nil {
		log.Fatalln(e)
	}
	e = dictionary.UpdateWord("name", "SSORCnyS")
	if e != nil {
		log.Fatalln(e)
	}
	word, e2 := dictionary.SearchWord("name")
	if e2 != nil {
		log.Fatalln(e)
	}
	fmt.Println(word)
}
