package main

import "strconv"

// Input: n = 15
// Output: ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]

func fizzBuzz(n int) []string {
	var result = make([]string, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				result[i-1] = "FizzBuzz"
				continue
			}
			result[i-1] = "Fizz"
			continue
		}
		if i%5 == 0 {
			result[i-1] = "Buzz"
			continue
		}
		result[i-1] = strconv.Itoa(i)
	}
	return result
}
