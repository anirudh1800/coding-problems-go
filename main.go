package main

import (
	"fmt"
)

func main() {

	testPalindrome()

}

func testPalindrome() {
	// Palindrome
	fmt.Println("Longest Palindrome")
	fmt.Println(longestPalindrome("babad"))

	testNthNodeLL()
}

func testSecondHighest() {
	// Second highest

	heap := createHeap(10)

	heap.addElement(5)
	heap.addElement(4)
	heap.addElement(3)
	heap.addElement(1)
	heap.addElement(6)

	fmt.Println(heap.length) // 5

	heap.heapify()

	fmt.Println(heap.popElement()) // 6

	fmt.Println(heap.length) // 4

	heap.heapify()

	fmt.Println(heap.popElement()) // 5

	fmt.Println(heap.length) // 3
}

func testNthNodeLL() {
	head := &ListNode{5, nil}

	head.Next = &ListNode{4, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{2, nil}
	head.Next.Next.Next.Next = &ListNode{1, nil}

	display(head)

	removeNthFromEnd(head, 4)

	display(head)

}
