package main

import (
	"fmt"
	"net/http"
)

func contentType(url string, outChannel chan string) {
	resp, err := http.Get(url)

	if err != nil {
		outChannel <- fmt.Sprintf("%s - %s", url, err)
	}

	defer resp.Body.Close() //

	ctype := resp.Header.Get("Content-Type")
	if ctype == "" {
		outChannel <- fmt.Sprintf("%s - %s", url, "Content-Type header is missing")
	}
	outChannel <- fmt.Sprintf("%s - %s", url, ctype)
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.lynda.com",
		"https://learning.oreilly.com",
		"https://stackoverflow.com",
		"https://golang.org",
	}
	ch := make(chan string)
	for _, url := range urls {
		go contentType(url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}
}
