/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

package main

// @lc code=start
func maxProfit(prices []int) int {
	left, right, maxProf := 0, 1, 0
	for right < len(prices) {
		if prices[right] > prices[left] {
			currentProf := prices[right] - prices[left]
			if currentProf > maxProf {
				maxProf = currentProf
			}
		} else {
			left = right
		}
		right++
	}
	return maxProf
}

// @lc code=end
