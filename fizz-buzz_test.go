package main

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		n      int
		result []string
	}{
		{
			n:      15,
			result: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}

	for _, c := range cases {
		if result := fizzBuzz(c.n); TestEq[string](result, c.result) != true {
			t.Errorf("Failure: Result %v, got %v for %v", c.result, result, c.n)
		} else {
			t.Logf("Success: Result %v, got %v for %v", c.result, result, c.n)
		}
	}
}
