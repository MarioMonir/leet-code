/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */
package main

// @lc code=start
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	mergedList := &ListNode{}
	ptr := mergedList

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			ptr.Next = list1
			list1 = list1.Next
		} else {
			ptr.Next = list2
			list2 = list2.Next
		}
		ptr = ptr.Next
	}

	if list1 != nil {
		ptr.Next = list1
	} else {
		ptr.Next = list2
	}

	return mergedList.Next
}

// @lc code=end
