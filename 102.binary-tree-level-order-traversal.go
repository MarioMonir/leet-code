/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 */

package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}
	var bfs func(node *TreeNode, level int)
	bfs = func(node *TreeNode, level int) {

		if len(res) <= level {
			res = append(res, []int{})
		}

		res[level] = append(res[level], node.Val)

		if node.Left != nil {
			bfs(node.Left, level+1)
		}
		if node.Right != nil {
			bfs(node.Right, level+1)
		}
	}

	bfs(root, 0)
	return res
}

// @lc code=end
