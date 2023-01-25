/*
 * @lc app=leetcode id=2011 lang=golang
 *
 * [2011] Final Value of Variable After Performing Operations
 */
package main

// @lc code=start
func finalValueAfterOperations(operations []string) int {
	count := 0
	for _, el := range operations {
		if el == "++X" || el == "X++" {
			count++
		} else {
			count--
		}
	}
	return count
}

// @lc code=end
