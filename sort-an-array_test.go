package main

import "testing"

func TestSortArray(t *testing.T) {
	cases := []struct {
		nums   []int
		result []int
	}{
		{
			nums:   []int{5, 2, 3, 1},
			result: []int{1, 2, 3, 5},
		},
		{
			nums:   []int{5, 1, 1, 2, 0, 0},
			result: []int{0, 0, 1, 1, 2, 5},
		},
		{
			nums:   []int{-2, 3, -5},
			result: []int{-5, -2, 3},
		},
	}

	for _, c := range cases {
		copyNumSlice := append([]int(nil), c.nums...)
		result := sortArray(copyNumSlice)
		if len(result) == len(c.result) {
			for i := 0; i < len(result); i++ {
				if result[i] == c.result[i] {
					continue
				}
				t.Errorf("Failure: Result %v, got %v for input %v", c.result, result, c.nums)
			}
			t.Logf("Success: Result %v, got %v for input %v", c.result, result, c.nums)
		} else {
			t.Errorf("Failure: Not same lenth. Result %v, got %v for input %v", c.result, result, c.nums)
		}
	}
}
