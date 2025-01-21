/*
 * @lc app=leetcode id=1920 lang=golang
 *
 * [1920] Build Array from Permutation
 */
package main

// @lc code=start
func buildArray(nums []int) []int {
	res := []int{}
	for i := range nums {
		res = append(res, nums[nums[i]])
	}
	return res
}

// @lc code=end
