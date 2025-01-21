/*
 * @lc app=leetcode id=589 lang=golang
 *
 * [589] N-ary Tree Preorder Traversal
 */
package main

type Node struct {
	Val      int
	Children []*Node
}

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	res := []int{}

	if root == nil {
		return []int{}
	}

	res = append(res, root.Val)

	var dfs func(node *Node)
	dfs = func(node *Node) {
		for _, n := range node.Children {
			res = append(res, n.Val)
			dfs(n)
		}
	}

	dfs(root)
	return res
}

// @lc code=end
