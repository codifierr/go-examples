package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func helloHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello gofers.")
}

// MathRequest func request
type MathRequest struct {
	Op    string  `json:"op"`
	Left  float64 `json:"left"`
	Right float64 `json:"right"`
}

// MathResponse func response
type MathResponse struct {
	Error  string  `json:"error"`
	Result float64 `json:"result"`
}

func mathHandler(rw http.ResponseWriter, r *http.Request) {
	//Decode request
	defer r.Body.Close()
	log.Printf("Request : %v", r)
	dec := json.NewDecoder(r.Body)
	req := &MathRequest{}

	if err := dec.Decode(req); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}
	// Do math
	resp := MathResponse{}
	switch req.Op {
	case "+":
		resp.Result = req.Left + req.Right
	case "-":
		resp.Result = req.Left - req.Right
	case "*":
		resp.Result = req.Left * req.Right
	case "/":
		if req.Right == 0.0 {
			resp.Error = "Divison by zero."
		} else {
			resp.Result = req.Left / req.Right
		}
	default:
		resp.Error = fmt.Sprintf("Unknown operation: %s", req.Op)
	}

	//Encode and return
	rw.Header().Set("Content-Type", "application/json")
	if resp.Error != "" {
		rw.WriteHeader(http.StatusBadRequest)
	}

	enc := json.NewEncoder(rw)
	if err := enc.Encode(resp); err != nil {
		log.Printf("Can't encode for client %v - %s", resp, err)
	}

}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/math", mathHandler)
	if err := http.ListenAndServe(":8390", nil); err != nil {
		log.Fatal(err)
	}

}
