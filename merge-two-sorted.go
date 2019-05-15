package main

import "fmt"

func main() {
	array1 := []int{1, 3, 7, 9, -1, -1, -1}
	array2 := []int{2, 4, 5}

	mergeTwoSortedArrays(&array1, &array2)

	fmt.Println(array1)

}

func mergeTwoSortedArrays(a *[]int, b *[]int) {
	i := len(*a) - 1
	for ; (*a)[i] < 0; i-- {
	}
	i++

	j := len(*b) - 1
	for ; i < len(*a) && j >= 0; i++ {
		(*a)[i] = (*b)[j]
		for k := i - 1; k >= 0 && (*a)[k] > (*a)[k+1]; k-- {
			temp := (*a)[k]
			(*a)[k] = (*a)[k+1]
			(*a)[k+1] = temp
		}
		j--
	}
}
