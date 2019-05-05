//making a http request
package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	//Get request
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalf("Error : can't call get method on provided url")
	}

	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
