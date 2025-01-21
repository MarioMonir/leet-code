package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	pointer := head

	for pointer.Next != nil {
		if pointer.Val == pointer.Next.Val  {
			pointer.Next = pointer.Next.Next
		}else{
			pointer = pointer.Next
		}
	}

	return head
}



