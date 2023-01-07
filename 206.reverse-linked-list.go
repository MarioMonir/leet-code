/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */

package main

// @lc code=start
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ptr1, ptr2 := head, head.Next
	for ptr2 != nil {
		temp := ptr1
		ptr1, ptr2 = ptr2, ptr2.Next
		ptr1.Next = temp
	}
	head.Next = nil
	return ptr1
}

// @lc code=end
