package main

/*
 * @lc app=leetcode id=1929 lang=golang
 *
 * [1929] Concatenation of Array
 */
// @lc code=start
func getConcatenation(nums []int) []int {
	n := len(nums)
	res := make([]int, 2*n)
	for i := range nums {
		res[i], res[i+n] = nums[i], nums[i]
	}
	return res
}

// @lc code=end
