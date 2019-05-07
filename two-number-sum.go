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
			result = make([]int, 2)
			if num < target - num {
			 result[0] = num
			 result[1] = target - num
			} else {
			 result[0] = target - num
			 result[1] = num
			}
		  break
		}
	}
	
	return result
}
