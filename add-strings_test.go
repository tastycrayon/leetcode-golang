package main

import (
	"testing"
)

func TestAddStrings(t *testing.T) {
	cases := []struct {
		num1   string
		num2   string
		result string
	}{
		{
			num1: "11", num2: "123",
			result: "134",
		},
		{
			num1: "456", num2: "77",
			result: "533",
		},
		{
			num1: "0", num2: "0",
			result: "0",
		},
	}

	for _, c := range cases {
		if result := addStrings(c.num1, c.num2); result != c.result {
			t.Errorf("Failure: Result %v, got %v for %v %v", c.result, result, c.num1, c.num2)
		} else {
			t.Logf("Success: Result %v, got %v for %v %v", c.result, result, c.num1, c.num2)
		}
	}
}
