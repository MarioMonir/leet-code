package main

func main() {

// Create a linked list
	listNode := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 3,
						Next: nil,
					},
				},
			},
		},
	}



	PrintNode(listNode)

	deleteDuplicates(listNode)

	PrintNode(listNode)
}
