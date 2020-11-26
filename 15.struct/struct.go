package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func structExample() {
	// * SynCROSS := person{"SynCROSS", 18}
	SynCROSS := person{name: "SynCROSS", age: 18}
	fmt.Println(SynCROSS.name, SynCROSS.age)
}
