/*
 * @lc app=leetcode id=724 lang=golang
 *
 * [724] Find Pivot Index
 */
package main

// @lc code=start
func pivotIndex(nums []int) int {
	var arrSum = func(arr []int) int {
		sum := 0
		for _, elem := range arr {
			sum += elem
		}
		return sum
	}
	left, right := 0, arrSum(nums)
	for i, elem := range nums {
		right -= elem
		if left == right {
			return i
		}
		left += elem
	}
	return -1
}

// @lc code=end
