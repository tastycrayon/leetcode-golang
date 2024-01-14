package main

import (
	"testing"
)

func TestMinWindow(t *testing.T) {
	cases := []struct {
		s      string
		t      string
		result string
	}{
		{
			s: "ADOBECODEBANC", t: "ABC",
			result: "BANC",
		},
		{
			s: "a", t: "a",
			result: "a",
		},
		{
			s: "a", t: "aa",
			result: "",
		},
		{
			s: "ahffaksfajeeubsne", t: "jefaa",
			result: "aksfaje",
		},
		{
			s: "aaffhkksemckelloe", t: "fhea",
			result: "affhkkse",
		},
	}

	for _, c := range cases {
		if result := minWindow(c.s, c.t); result != c.result {
			t.Errorf("Failure: Result %v, got %v for %v %v", c.result, result, c.s, c.t)
		} else {
			t.Logf("Success: Result %v, got %v for %v %v", c.result, result, c.s, c.t)
		}
	}
}
