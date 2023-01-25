/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */
package main

// @lc code=start
func removeElement(nums []int, val int) int {
	count := 0
	for _, e := range nums {
		if val != e {
			nums[count] = e
			count++
		}
	}
	return count
}

// @lc code=end
