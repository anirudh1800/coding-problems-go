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
