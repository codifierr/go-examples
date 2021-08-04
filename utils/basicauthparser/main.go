package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/test", helloHandler)

	if err := http.ListenAndServe(":8190", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(rw http.ResponseWriter, r *http.Request) {
	username, pass, ok := parseUserAuthorizationHeader(r)
	if ok {
		fmt.Println("username", username)
		fmt.Println("pass", pass)
	}
	fmt.Fprint(rw, "Hello gofers.")

}

func parseUserAuthorizationHeader(r *http.Request) (username, password string, ok bool) {
	auth := r.Header.Get("UserAuthorization")
	if auth == "" {
		return
	}
	const prefix = "Basic "
	if len(auth) < len(prefix) || !strings.EqualFold(auth[:len(prefix)], prefix) {
		return
	}
	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return
	}
	cs := string(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return
	}
	return cs[:s], cs[s+1:], true
}
