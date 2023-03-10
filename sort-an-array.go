package main

func merge(nums []int, start, mid, end int) {
	L := make([]int, (mid + 1 - start)) //start 0, mid 1 and 1-0 = 1, so 2-0 = 2
	copy(L, nums[start:mid+1])          // mid + 1 mean start 0, mid 2

	R := make([]int, (end - mid))
	copy(R, nums[mid+1:end+1])

	i, j, k := 0, 0, start

	for i < len(L) && j < len(R) {
		if L[i] <= R[j] {
			nums[k] = L[i]
			i++
		} else {
			nums[k] = R[j]
			j++
		}
		k++
	}
	for i < len(L) {
		nums[k] = L[i]
		i++
		k++
	}
	for j < len(R) {
		nums[k] = R[j]
		j++
		k++
	}
}
func mergeSort(nums []int, start int, end int) {
	if end-start+1 <= 1 {
		return
	}
	// split half
	mid := (start + end) / 2

	mergeSort(nums, start, mid)
	mergeSort(nums, mid+1, end)

	merge(nums, start, mid, end)

}
func sortArray(nums []int) []int {
	mergeSort(nums, 0, len(nums)-1)
	return nums
}
