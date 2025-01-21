/*
 * @lc app=leetcode id=1512 lang=golang
 *
 * [1512] Number of Good Pairs
 */

package main

// @lc code=start
func numIdenticalPairs(nums []int) int {
	count, hashMap := 0, map[int]int{}
	for _, el := range nums {
		count += hashMap[el]
		hashMap[el]++
	}
	return count
}

// @lc code=end
