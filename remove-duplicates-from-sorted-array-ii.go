package main

// Input: nums = [1,1,1,2,2,3]
// Output: 5, nums = [1,1,2,2,3,_]
func removeDuplicates(nums []int) int {
	var right, left int = 0, 0
	for right < len(nums) {
		count := 1
		for right+1 < len(nums) && nums[right] == nums[right+1] {
			right++
			count++
		}
		if count > 2 {
			count = 2 // cap it to 2
		}
		for i := 0; i < count; i++ {
			nums[left] = nums[right]
			left++
		}
		right++
	}

	return left
}
