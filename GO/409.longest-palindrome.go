package main

/*
 * @lc app=leetcode id=409 lang=golang
 *
 * [409] Longest Palindrome
 */
// @lc code=start
func longestPalindrome(s string) int {
	count, charCount := 0, map[string]int{}
	for _, ch := range s {
		charCount[string(ch)]++
		if charCount[string(ch)]%2 == 0 {
			count += 2
		}
	}
	if len(s) > count {
		count += 1
	}
	return count
}

// @lc code=end
