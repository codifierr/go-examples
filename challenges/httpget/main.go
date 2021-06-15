package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close() //

	ctype := resp.Header.Get("Content-Type")
	if ctype == "" {
		return "", fmt.Errorf("Content-Type header is missing")
	}

	return ctype, nil
}

func main() {
	url := "https://www.google.com"
	ctype, err := contentType(url)

	if err != nil {
		fmt.Printf("Error : %s\n", err)
	} else {
		fmt.Println(ctype)
	}

}
