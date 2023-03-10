package main

import (
	"sort"
	"sync"
)

func missingNumber(nums []int) int {
	numberOfItems := len(nums)
	sumOfRange := numberOfItems * (numberOfItems + 1) / 2

	for _, item := range nums {
		sumOfRange -= item
	}
	return sumOfRange
}

func missingNumber1(nums []int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	numberOfItems := len(nums)
	if numberOfItems == 1 {
		if nums[0] == 0 {
			return 1
		} else {
			return 0
		}
	}
	if (nums[numberOfItems-1]) != numberOfItems {
		return numberOfItems
	}
	for i := 1; i < numberOfItems; i++ {
		if nums[i]-nums[i-1] == 1 {
			continue
		}
		return (nums[i] + nums[i-1]) / 2
	}
	return 0
}

// concurrent
func missingNumber2(nums []int) int {
	var wg sync.WaitGroup = sync.WaitGroup{}

	numberOfItems := len(nums)
	sumOfRange := (numberOfItems * (numberOfItems + 1)) / 2

	result := make(chan int, 1)
	defer close(result)

	result <- sumOfRange // init value

	for _, item := range nums {
		wg.Add(1)
		current := <-result
		go func(item int) {
			result <- subtract(current, item) // continue subtracting
			wg.Done()
		}(item)
	}
	wg.Wait()
	return <-result
}
func subtract(a, b int) int {
	return a - b
}
