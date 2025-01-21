/*
 * @lc app=leetcode id=278 lang=golang
 *
 * [278] First Bad Version
 */
package main

// @lc code=start
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	r := n
	for l := 1; l < r; {
		m := l + (r-l)/2
		if false {
			r = m
		} else {
			l = m + 1
		}
	}
	return r
}

// @lc code=end
