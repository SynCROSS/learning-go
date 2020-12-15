package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request is Failed.")

func main() {
	var results map[string]string

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
		// ! Errors may occur.
		// ! Because empty map without braces can NOT be written
		results[url] = url
		hitURL(url)
	}

}

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	res, err := http.Get(url)
	if err == nil || res.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
