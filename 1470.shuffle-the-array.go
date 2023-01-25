/*
 * @lc app=leetcode id=1470 lang=golang
 *
 * [1470] Shuffle the Array
 */
package main

// @lc code=start
func shuffle(nums []int, n int) []int {
	res := []int{}
	i, j := 0, n
	for len(res) != len(nums) {
		res = append(res, nums[i])
		res = append(res, nums[j])
		j, i = j+1, i+1
	}
	return res
}

// @lc code=end
