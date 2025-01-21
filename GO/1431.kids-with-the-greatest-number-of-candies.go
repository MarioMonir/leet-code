/*
 * @lc app=leetcode id=1431 lang=golang
 *
 * [1431] Kids With the Greatest Number of Candies
 */

package main

// @lc code=start
func kidsWithCandies(candies []int, extraCandies int) []bool {
	max, res := 0, []bool{}
	for _, c := range candies {
		if c > max {
			max = c
		}
	}
	for _, c := range candies {
		if c+extraCandies >= max {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}
	return res
}

// @lc code=end
