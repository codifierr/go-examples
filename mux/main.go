package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)

	go func() {
		fmt.Println("listen on 8080")
		log.Fatal(http.ListenAndServe(":8080", router))
	}()

	fmt.Println("listen on 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
