package main

import "fmt"

// NewSubscriptionsMap function to create sunscription map
func NewSubscriptionsMap() map[string]bool {
	return map[string]bool{
		"test":     false,
		"test1":          false,
		"test2":       false,
	}
}

func main() {
	testmap := NewSubscriptionsMap()
	if _, ok := testmap["fo"]; ok {
		//fmt.Println(val)
	}
	fmt.Println(testmap)
}
