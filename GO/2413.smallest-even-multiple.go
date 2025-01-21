/*
 * @lc app=leetcode id=2413 lang=golang
 *
 * [2413] Smallest Even Multiple
 */

package main

// @lc code=start
func smallestEvenMultiple(n int) int {
    if n % 2 == 0 {
        return n
    } 
    return 2 * n 
}
// @lc code=end

