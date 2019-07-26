package program

type BST struct {
	value int
	left  *BST
	right *BST
}

func (tree *BST) FindClosestValue(target int) int {
	var rootDiff int 
	
	rootDiff = Abs(target - tree.value)
	var leftNo int
	var rightNo int 
	
	leftNo = -1 
	if tree.value > target && tree.left != nil {
		leftNo = tree.left.FindClosestValue(target) 			
	}
	
	rightNo = -1
	if tree.value < target && tree.right != nil {
		rightNo = tree.right.FindClosestValue(target)
	}
	
	if leftNo != -1 && Abs(leftNo - target) < rootDiff {
		 return leftNo
	} else if rightNo != -1 && Abs(rightNo - target) < rootDiff {
		return rightNo
	} else {
		return tree.value
	}
}

func Min(x int, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

