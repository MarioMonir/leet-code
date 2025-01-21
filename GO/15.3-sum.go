/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */
package main

import "sort"

// @lc code=start
func threeSum(nums []int) [][]int {

	if len(nums) < 3 {
		return [][]int{}
	}

	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-2; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		front := i + 1
		rear := len(nums) - 1

		for front < rear {
			sum := nums[i] + nums[front] + nums[rear]
			if sum > 0 {
				rear--
			} else if sum < 0 {
				front++
			} else {
				res = append(res, []int{nums[i], nums[front], nums[rear]})
				front++
				for front < rear && nums[front] == nums[front-1] {
					front++
				}
			}
		}
	}
	return res
}

// @lc code=end
