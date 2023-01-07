/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

package main

// @lc code=start
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sums := &ListNode{}
	res := sums
	var carrier = 0

	for l1 != nil || l2 != nil || carrier != 0 {

		var v1 int = 0
		var v2 int = 0

		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := v1 + v2 + carrier
		carrier = int(sum / 10)
		unit := sum % 10
		sums.Next = &ListNode{Val: unit}

		sums = sums.Next
	}

	return res.Next

}

// @lc code=end
