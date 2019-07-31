package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type healthcheckresults struct {
	Results []result `json:"results"`
}
type result struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Reason string `json:"reason,omitempty"`
}

func main() {
	result1 := result{
		Name:   "test1",
		Status: "Ok",
		Reason: "Great",
	}
	result2 := result{
		Name:   "test2",
		Status: "Ok",
		Reason: "Great",
	}

	checkResults := healthcheckresults{
		Results: []result{result1, result2},
	}
	b, err := json.Marshal(checkResults)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}
