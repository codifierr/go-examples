package main

import (
	"log"
	"math/cmplx"
)

var (
	toBe          = false
	maxInt uint64 = 1<<64 - 1
	z             = cmplx.Sqrt(-5 + 12i)
)

// main is the entry point of the program.
func main() {
	printVariable("toBe", toBe)
	printVariable("maxInt", maxInt)
	printVariable("z", z)
}

// printVariable prints the type and value of a variable.
func printVariable(name string, value interface{}) {
	log.Printf("Variable %s is %v\n", name, value)
}
