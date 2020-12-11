package main

import (
	"fmt"
	"log"
	"main/dict"
)

func main() {
	dictionary := dict.Dictionary{"name": "SynCROSS"}
	definition, e := dictionary.Search("name")

	if e != nil {
		log.Fatalln(e)
	}
	fmt.Println(definition)
}
