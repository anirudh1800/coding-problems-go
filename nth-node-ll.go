package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // ListNode in LL
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slowPtr := head
	fastPtr := head
	prevNode := head

	for ; n > 0; n-- {
		fastPtr = fastPtr.Next
	}

	for fastPtr != nil {
		prevNode = slowPtr
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next
	}

	prevNode.Next = slowPtr.Next

	if slowPtr == head {
		head = head.Next
	}

	return slowPtr

}

func display(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}

	fmt.Println()
}
