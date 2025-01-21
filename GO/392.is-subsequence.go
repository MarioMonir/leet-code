/*
 * @lc app=leetcode id=392 lang=golang
 *
 * [392] Is Subsequence
 */

package main

// @lc code=start
func isSubsequence(s string, t string) bool {
	ptrS, ptrT := 0, 0
	for ptrS < len(s) && ptrT < len(t) {
		if s[ptrS] == t[ptrT] {
			ptrS++
		}
		ptrT++
	}
	return ptrS == len(s)
}

// @lc code=end
