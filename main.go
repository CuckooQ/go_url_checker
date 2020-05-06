package main

import (
	"errors"
	"fmt"
	"net/http"
)

// error
var errRequestFailed error = errors.New("Request failed")

// hit url
func hitURL(url string) error {
	fmt.Println("Checking URL...", url)
	resp, err := http.Get(url)

	if err != nil || resp.StatusCode >= 400 {
		return err
	}

	return nil
}

// set results
func setResultsFromUrl(url []string, results map[string]string) map[string]string {
	result := "OK"

	for _, url := range url {
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}

		results[url] = result
	}

	return results
}

// print results
func printResults(results map[string]string) {
	for key, value := range results {
		fmt.Println(key, value)
	}
}

func main() {
	results := make(map[string]string)
	url := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	// check url
	results = setResultsFromUrl(url, results)

	// print results
	printResults(results)
}
