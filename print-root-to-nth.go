package main

import "fmt"
import "github.com/emirpasic/gods/stacks/arraystack"

type treeNode struct {
	Data  int
	Left  *treeNode
	Right *treeNode
}

func main() {
	root := &treeNode{4, nil, nil}

	root.Left = &treeNode{2, nil, nil}
	root.Right = &treeNode{3, nil, nil}

	root.Left.Left = &treeNode{1, nil, nil}

	// printRootToNthNode(root, 1)
	// fmt.Println(root.Data)

	stack := arraystack.New()
	printRootToNthNode2(root, 1, stack)

	fmt.Println(stack.Values())
}

func printRootToNthNode2(root *treeNode, n int, target *arraystack.Stack) bool {
	if root == nil {
		return false
	}

	target.Push(root.Data)
	if root.Data == n {
		return true
	}

	if printRootToNthNode2(root.Left, n, target) == true || 
		printRootToNthNode2(root.Right, n, target) == true {
		return true
	}

	target.Pop()

	return false
}

func printRootToNthNode(root *treeNode, n int) int {
	if root.Data == n {
		return 1
	}

	if root.Left != nil {
		leftResult := printRootToNthNode(root.Left, n)
		if leftResult == 1 {
			fmt.Println(root.Left.Data)
			return 1
		}
	}

	if root.Right != nil {
		rightResult := printRootToNthNode(root.Right, n)
		fmt.Println(root.Right.Data)
		if rightResult == 1 {
			return 1
		}
	}

	return 0
}
