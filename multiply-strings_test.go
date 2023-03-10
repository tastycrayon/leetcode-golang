package main

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	cases := []struct {
		num1   string
		num2   string
		result string
	}{
		{
			num1: "2", num2: "3",
			result: "6",
		},
		{
			num1: "123", num2: "456",
			result: "56088",
		},
		{
			num1: "9", num2: "9",
			result: "81",
		},
		{
			num1: "925101087184894", num2: "3896737933784656127",
			result: "3604876499018802870538090258945538",
		},
	}

	for _, c := range cases {
		if result := multiply(c.num1, c.num2); result != c.result {
			t.Errorf("Failure: Result %v, got %v for %v %v", c.result, result, c.num1, c.num2)
		} else {
			t.Logf("Success: Result %v, got %v for %v %v", c.result, result, c.num1, c.num2)
		}
	}
}
