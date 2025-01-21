/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
 */
package main

import (
	"math"
)

// @lc code=start
func isValidBST(root *TreeNode) bool {
	// n: node , l: lowest,  h: highest
	var validate func(n *TreeNode, l int, h int) bool
	validate = func(n *TreeNode, l int, h int) bool {
		if n == nil {
			return true
		}
		if l >= n.Val || n.Val >= h {
			return false
		}
		return validate(n.Left, l, n.Val) && validate(n.Right, n.Val, h)
	}
	return validate(root, math.MinInt, math.MaxInt)
}

// @lc code=end
