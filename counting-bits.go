package main

import "math/bits"

func count(n int) (result int) {
	for n != 0 {
		n = n & (n - 1)
		result++
	}
	return result
}

func countBits(n int) []int {
	var result = make([]int, 0, n+1)
	for i := 0; i < n+1; i++ {
		result = append(result, count(i))
	}
	return result
}
func countBits1(n int) []int {
	var result = make([]int, 0, n+1)
	for i := 0; i < n+1; i++ {
		result = append(result, bits.OnesCount(uint(i)))
	}
	return result
}
