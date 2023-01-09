/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
 */
package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	var validate func(n *TreeNode, min int, max int) bool
	validate = func(n *TreeNode, min int, max int) bool {
		if n == nil {
			return true
		}
		if min >= n.Val || max <= n.Val {
			return false
		}
		return validate(n.Left, min, n.Val) && validate(n.Right, n.Val, max)
	}
	return validate(root, math.MinInt, math.MaxInt)
}

// @lc code=end
