package main

import "fmt"

// NewSubscriptionsMap function to create sunscription map
func NewSubscriptionsMap() map[string]bool {
	return map[string]bool{
		"monitoring":     false,
		"apprf":          false,
		"presence":       false,
		"cloudguest":     false,
		"usermanagement": false,
		"visualrf":       false,
		"contextengine":  false,
		"ucc":            false,
		"lync":           false,
		"audit":          false,
		"clickstream":    false,
		"location":       false,
		"rapids":         false,
		"syslog":         false,
		"stream-stats":   false,
	}
}

func main() {
	testmap := NewSubscriptionsMap()
	if _, ok := testmap["fo"]; ok {
		//fmt.Println(val)
	}
	fmt.Println(testmap)
}
