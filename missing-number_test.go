package main

import "testing"

func TestMissingNumber(t *testing.T) {
	cases := []struct {
		nums   []int
		result int
	}{
		{
			nums:   []int{3, 0, 1},
			result: 2,
		},
		{
			nums:   []int{0, 1},
			result: 2,
		},
		{
			nums:   []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			result: 8,
		},
	}

	for _, c := range cases {
		if result := missingNumber(c.nums); result != c.result {
			t.Errorf("Failure: Result %v, got %v for %v", c.result, result, c.nums)
		} else {
			t.Logf("Success: Result %v, got %v for %v", c.result, result, c.nums)
		}

		if result := missingNumber1(c.nums); result != c.result {
			t.Errorf("Failure: Result %v, got %v for %v", c.result, result, c.nums)
		} else {
			t.Logf("Success: Result %v, got %v for %v", c.result, result, c.nums)
		}

		if result := missingNumber2(c.nums); result != c.result {
			t.Errorf("Failure: Result %v, got %v for %v", c.result, result, c.nums)
		} else {
			t.Logf("Success: Result %v, got %v for %v", c.result, result, c.nums)
		}
	}
}
