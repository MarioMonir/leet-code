/*
 * @lc app=leetcode id=1672 lang=golang
 *
 * [1672] Richest Customer Wealth
 */
package main

// @lc code=start
func maximumWealth(accounts [][]int) int {
    max := 0
    for _, c := range accounts{
        w := 0 
        for _, a := range c {
            w += a 
        }
        if w > max {
            max = w
        }
    } 
    return max
}
// @lc code=end

