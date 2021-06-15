package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var data = `{
	"user" : "satyendra",
	"type":"deposit",
	"amount":100.00
}`

//Request for bank operation's
type Request struct {
	Login  string  `json:"user"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

func main() {
	rdr := bytes.NewBufferString(data) // simulate a file socket

	//decode request
	dec := json.NewDecoder(rdr)

	req := &Request{}

	if err := dec.Decode(req); err != nil {
		log.Fatalf("Error can't decode - %s\n", err)
	}

	fmt.Printf("Decoded string - %+v\n", req)

	//create response
	prevbalance := 45000.00
	resp := map[string]interface{}{
		"ok":      true,
		"balance": prevbalance + req.Amount,
	}

	//encode response
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(resp); err != nil {
		log.Fatalf("Error can't encode - %s\n", err)
	}

	//fmt.Printf("Encoded response - %+v\n",resp)
}
