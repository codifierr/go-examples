package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		name string
	}{
		// Testing the printVariable function with different variables.
		{name: "toBe"},
		{name: "maxInt"},
		{name: "z"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Redirecting the output of the printVariable function to a buffer.
			var outputBuffer bytes.Buffer
			log.SetOutput(&outputBuffer)

			// Calling the printVariable function with the current test case.
			printVariable(tt.name, getVariable(tt.name))

			// Getting the output from the buffer.
			output := outputBuffer.String()

			splitStr := strings.SplitN(output, " ", 3)
			var strWithoutTimestamp string
			if len(splitStr) > 1 {
				strWithoutTimestamp = splitStr[2]
			} else {
				log.Println("Invalid input string")
			}

			// Verifying the expected output.
			expectedOutput := fmt.Sprintf("Variable %s is %v\n", tt.name, getVariable(tt.name))
			if strWithoutTimestamp != expectedOutput {
				t.Errorf("Unexpected output. Expected: %s, got: %s", expectedOutput, output)
			}
		})
	}
}

// getVariable returns the value of the given variable.
func getVariable(name string) interface{} {
	switch name {
	case "toBe":
		return toBe
	case "maxInt":
		return maxInt
	case "z":
		return z
	default:
		return nil
	}
}
