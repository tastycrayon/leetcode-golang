package main

import "testing"

func TestFindTheDifference(t *testing.T) {
	cases := []struct {
		s      string
		t      string
		result byte
	}{
		{
			s:      "abcd",
			t:      "abcde",
			result: byte('e'),
		},
		{
			s:      "",
			t:      "e",
			result: byte('e'),
		},
	}

	for _, c := range cases {
		if result := findTheDifference(c.s, c.t); result != c.result {
			t.Errorf("Failure: Result %v, got %v for %v, %v", c.result, result, c.s, c.t)
		} else {
			t.Logf("Success: Result %v, got %v for %v, %v", c.result, result, c.s, c.t)
		}
		if result := findTheDifference1(c.s, c.t); result != c.result {
			t.Errorf("Failure: Result %v, got %v for %v, %v", c.result, result, c.s, c.t)
		} else {
			t.Logf("Success: Result %v, got %v for %v, %v", c.result, result, c.s, c.t)
		}

	}
}
