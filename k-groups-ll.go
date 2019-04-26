package main

// 1-> 2 -> 3 -> 4 -> 5
// 2 -> 1 -> 4 -> 3 -> 5

func reverseLL(head *ListNode, k int) *ListNode {
	current := head
	var prev *ListNode
	var next *ListNode
	count := 0
	for current != nil && count < k {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
		count++
	}

	if next != nil {
		head.Next = reverseLL(next, k)
	}

	return prev
}
