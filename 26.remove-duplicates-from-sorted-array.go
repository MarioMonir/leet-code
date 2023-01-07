/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */
package main

// @lc code=start
func removeDuplicates(nums []int) int {
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			nums[k] = nums[i]
			k += 1
		}
	}
	return k
}

// @lc code=end
