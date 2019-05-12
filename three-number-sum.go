package main

import "fmt"
import "sort"
import "reflect"

func main() {

	result := threeNumberSum([]int{1, 2,  4, 5}, 7)

	fmt.Println(result)

	fmt.Println(reflect.DeepEqual([][]int{{1, 2, 4}}, result))

}

func threeNumberSum(array []int, target int) [][]int {

	result := [][]int{}

	sort.Ints(array)

	for i := 0; i < len(array) - 2; i++ {
		left := i + 1
		right := len(array) - 1
		for left < right {
			currentSum := array[i] + array[left] + array[right]
			if target == currentSum {
				result = append(result, []int{array[i], array[left], array[right]})
				left = left + 1
				right = right - 1
			} else if currentSum > target {
				right = right - 1
			} else if currentSum < target {
				left = left + 1
			}
		}
	}

	return result
}

func twoNumberSum(array []int, target int) []int {
	result := []int{}
	hashMap := make(map[int]int)

	for idx, num := range array {
		hashMap[num] = idx
	}

	for idx1, num := range array {

		idx2, present := hashMap[target-num]

		if present && idx1 != idx2 {
			if num < target-num {
				result = []int{num, target - num}
			} else {
				result = []int{target - num, num}
			}
			break
		}
	}

	return result

}
