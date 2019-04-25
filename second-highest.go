package main

import (
	"fmt"
)

// Find 2nd highest element in an array

// Heap struct to hold data
type Heap struct {
	items  []int
	length int
}

func createHeap(length int) *Heap {
	items := make([]int, length)
	return &Heap{items, 0}
}

func (h *Heap) addElement(key int) bool {
	if h.length < len(h.items) {
		h.items[h.length] = key
		h.length++
	} else {
		fmt.Println("Heap size exceeded")
		return false
	}
	return true
}

func (h *Heap) peekElement() int {
	return h.items[0]
}

func (h *Heap) popElement() int {
	item := h.items[0]
	h.items[0] = h.items[h.length-1]
	h.length = h.length - 1
	return item
}

func (h *Heap) heapify() {
	n := h.length / 2
	for ; n > 0; n = n / 2 {

		if n-1 >= 0 {
			h.maxHeap(n - 1)
		}

		if n-2 >= 0 {
			h.maxHeap(n - 2)
		}
	}
}

func (h *Heap) maxHeap(root int) {

	i := root

	if h.length > (root*2+1) && h.items[root*2+1] > h.items[i] {
		i = root*2 + 1
	}

	if h.length > (root*2+2) && h.items[root*2+2] > h.items[i] {
		i = root*2 + 2
		fmt.Println(i)

	}

	if i != root {
		temp := h.items[i]
		h.items[i] = h.items[root]
		h.items[root] = temp
	}
}
