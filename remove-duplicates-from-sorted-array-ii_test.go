package main

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	cases := []struct {
		nums       []int
		resultNums []int
		result     int
	}{
		{
			nums:       []int{1, 1, 1, 2, 2, 3},
			resultNums: []int{1, 1, 2, 2, 3},
			result:     5,
		},
	}

	for _, c := range cases {
		originalNum := make([]int, len(c.nums))
		copy(originalNum, c.nums)
		if result := removeDuplicates(c.nums); TestEq[int](c.resultNums, c.nums[:c.result]) != true {
			t.Errorf("Failure: Result %v and %v, got %v and %v for %v", c.result, c.resultNums, result, c.nums, originalNum)
		} else {
			t.Logf("Success: Result %v and %v, got %v and %v for %v", c.result, c.resultNums, result, c.nums, originalNum)
		}
	}
}
