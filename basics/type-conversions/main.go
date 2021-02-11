package main

import (
	"fmt"
	"reflect"
	"strconv"
	//"strconv"
)

func main() {
	// var x, y int = 4, 5
	// var f = math.Sqrt(float64(x*x + y*y))
	// var z = uint(f)
	// fmt.Println(x, y, z)
	// fmt.Println("Hello, playground")

	var tst interface{} = "satyendra"
	var tst2 interface{} = 5000232
	tst3 := 1.2
	var customerID string
	i := 1
	val := reflect.ValueOf(i)
	fmt.Println(!val.IsZero())

	fmt.Println(reflect.TypeOf(tst))
	fmt.Println(reflect.TypeOf(tst2))
	fmt.Println(reflect.TypeOf(tst3))
	customerID, ok := tst2.(string)
	fmt.Println(customerID)
	fmt.Println(ok)
	if !ok {
		fmt.Println("convert")
		customerID = strconv.Itoa(tst2.(int))
	}
	fmt.Println(customerID)
}
