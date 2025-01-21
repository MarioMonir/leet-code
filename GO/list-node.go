package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintNode(head *ListNode) {
	pointer := head
	for pointer != nil {
		fmt.Print(pointer.Val , " ,")
		pointer = pointer.Next
	}
	fmt.Println()
}