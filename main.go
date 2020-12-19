package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request is Failed")

func main() {
	// * var results = map[string]string{}
	var results = make(map[string]string) // * is same if you use 'make' function.

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.uber.com/",
		"https://github.com/",
		"https://soundcloud.com/",
		"https://www.fb.com/",
		"https://www.instagram.com/",
		"https://www.udemy.com/",
	}

	for _, url := range urls {
		result := "Success"
		e := hitURL(url)
		if e != nil {
			result = "Fail"
		}
		results[url] = result
	}
	fmt.Println()
	fmt.Println("-------------------------")
	fmt.Println()
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	res, err := http.Get(url)
	// * I'm sorry. There was a typo.
	// * err == nil => err != nil
	if err != nil || res.StatusCode >= 400 {
		fmt.Println(err, res.StatusCode)
		return errRequestFailed
	}
	return nil
}
