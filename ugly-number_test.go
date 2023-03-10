package main

import (
	"testing"
)

func TestIsUgly(t *testing.T) {
	cases := []struct {
		num    int
		result bool
	}{
		{
			num:    6,
			result: true,
		},
		{
			num:    1,
			result: true,
		},
		{
			num:    14,
			result: false,
		},
	}

	for _, c := range cases {
		if result := isUgly(c.num); result != c.result {
			t.Errorf("Failure: Result %v, got %v for %v", c.result, result, c.num)
		} else {
			t.Logf("Success: Result %v, got %v for %v", c.result, result, c.num)
		}
	}
}
