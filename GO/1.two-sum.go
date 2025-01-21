/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
package main

// @lc code=start
func twoSum(nums []int, target int) []int {
	hash := map[int]int{}
	for i, el := range nums {
		if _, ok := hash[el]; ok {
			return []int{hash[el], i}
		}
		hash[target-el] = i
	}
	return nil
}

// @lc code=end
