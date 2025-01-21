/*
 * @lc app=leetcode id=2469 lang=golang
 *
 * [2469] Convert the Temperature
 */
package main

// @lc code=start
func convertTemperature(celsius float64) []float64 {
	return []float64{celsius + 273.15, (celsius * 1.80) + 32.00}
}

// @lc code=end
