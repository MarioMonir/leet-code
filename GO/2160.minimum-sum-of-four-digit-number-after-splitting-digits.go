/*
 * @lc app=leetcode id=2160 lang=golang
 *
 * [2160] Minimum Sum of Four Digit Number After Splitting Digits
 */

package main

import "sort"

// @lc code=start
func minimumSum(num int) int {
	nums := []int{0, 0, 0, 0}
	for i := 0; i < 4; i++ {
		nums[i] = num % 10
		num /= 10
	}
	sort.Ints(nums)
	return (nums[0]*10 + (nums[3])) + (nums[1]*10 + (nums[2]))
}

// @lc code=end
