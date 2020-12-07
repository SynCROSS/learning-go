package main

import (
	"fmt"
)

func main() {
	object := map[string]string{"name": "SynCROSS", "age": "18"} // * map[type of key]type of value
	for _, value := range object {
		fmt.Println(value)
	}
}
