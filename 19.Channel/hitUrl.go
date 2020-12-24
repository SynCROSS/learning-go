package main

import (
	"errors"
	"fmt"
	"net/http"
)

type hitResult struct {
	url              string
	connectionResult string
}

var errRequestFailed = errors.New("Request is Failed")

func main() {
	requestResults := make(map[string]string)
	// * When you use structure as a parameter for the Channel,
	// * you must write a name not a format.
	channel := make(chan hitResult)

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
		go hitURL(url, channel)
	}

	for i := 0; i < len(urls); i++ {
		result := <-channel
		requestResults[result.url] = result.connectionResult
	}

	for url, connectionResult := range requestResults {
		fmt.Println(url + " | " + connectionResult)
	}

}

func hitURL(url string, channel chan<- hitResult) { // * chan<- means that Channel is Send-Only in this function
	res, err := http.Get(url)
	result := "Success"
	if err != nil || res.StatusCode >= 400 {
		result = "Failure"
	}
	channel <- hitResult{url: url, connectionResult: result}
}
