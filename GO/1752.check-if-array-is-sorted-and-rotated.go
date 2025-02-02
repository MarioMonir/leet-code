package main

// the brute force
// check that every element is less than the next element
// and this rule can be only break once
func check(nums []int) bool {
	allowedBreaks := 0

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			allowedBreaks++

			if allowedBreaks > 1 {
				return false
			}
		}

	}

	if allowedBreaks == 1 && nums[0] < nums[len(nums)-1] {
		return false
	}

	return true
}
