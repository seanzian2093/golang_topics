// +build main

package main

import (
	"testing"
)

// go test or go test -v in command line to see testing results
func TestCalculate(t *testing.T) {
	if Calculate(2) !=4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}

func TestTableCalculate(t *testing.T) {
	// tests is a slice of struct of
	// 2 fields and initilize at
	// {}
	var tests = []struct {
		input int
		expected int
	} {
		{2, 4},
		{-1, 1},
		{0, 2},
		{9999, 10001},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test Failed: {} input, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
