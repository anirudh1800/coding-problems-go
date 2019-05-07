package main

import "fmt"

func main() {
	output := twoNumberSum([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10)

	fmt.Println(output)
}

func twoNumberSum(array []int, target int) []int {
	result := []int{}
	hashMap := make(map[int]int)
	
	for idx, num := range array {
		hashMap[num] = idx
	}
	
	for idx1, num := range array {
		
		idx2, present := hashMap[target - num]
	
		if present && idx1 != idx2 {
			if num < target - num {
				result = []int{num , target - num }
			} else {
				result = []int{target - num , num }
			}
		  break
		}
	}
	
	return result
}
