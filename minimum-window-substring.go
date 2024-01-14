package main

import (
	"math"
)

func minWindow(s string, t string) string {
	var (
		table       []int = make([]int, 128)
		window      []int = make([]int, 128)
		have        int   = 0
		need        int   = len(t)
		leftIdx     int   = 0
		resultStart       = 0
		minimum           = math.MaxInt
	)

	for _, v := range t {
		table[v]++
	}

	for rightIdx, rightChar := range s {
		window[rightChar]++

		// is in table, and same size as window
		if acc := table[rightChar]; acc > 0 && acc >= window[rightChar] {
			have++
		}
		for need == have {
			if m := rightIdx - leftIdx; m < minimum {
				minimum = m
				resultStart = leftIdx
			}
			// pop from left of window
			leftChar := s[leftIdx]
			window[leftChar]--
			if acc := table[leftChar]; acc > 0 && window[leftChar] < acc {
				have--
			}
			leftIdx++
		}
	}
	if minimum == math.MaxInt {
		return ""
	}
	return s[resultStart:(resultStart + minimum + 1)]
}

// Input: s = "ADOBECODEBANC", t = "ABC"
// Output: "BANC"
// Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.

// func main() {
// 	println(minWindow("ADOBECODEBANC", "ABC")) // 1
// }
