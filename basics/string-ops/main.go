package main

import (
	"fmt"
	"strings"
)

func main() {

	s := strings.Split("central:^:c437a56b4a554c509a84b1f0adf1d5f5:^:satyendra.singh3339+2@gmail.com", ":")
	fmt.Println(s[1])
}
