package main

import (
	"testing"
)

func TestCountBits(t *testing.T) {

	cases := []struct {
		n      int
		result []int
	}{
		{n: 2, result: []int{0, 1, 1}},
		{n: 5, result: []int{0, 1, 1, 2, 1, 2}},
	}
	for _, c := range cases {
		if result := countBits(c.n); TestEq(result, c.result) != true {
			t.Errorf("Failure: Result %v, got %v for %v", c.result, result, c.n)
		} else {
			t.Logf("Success: Result %v, got %v for %v", c.result, result, c.n)
		}
		if result := countBits1(c.n); TestEq(result, c.result) != true {
			t.Errorf("Failure: Result %v, got %v for %v", c.result, result, c.n)
		} else {
			t.Logf("Success: Result %v, got %v for %v", c.result, result, c.n)
		}
	}
}
