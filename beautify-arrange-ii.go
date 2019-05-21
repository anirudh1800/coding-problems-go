package main

import "fmt"

func main() {
	fmt.Println(constructArray(10, 9))
}

func constructArray(n int, k int) []int {

	ans := make([]int, n)
	c := 0

	for v := 1; v < n-k; v++ {
		ans[c] = v
		c++
	}

	for v := 0; v <= k; v++ {
		if v%2 == 0 {
			ans[c] = n - k + v/2
		} else {
			ans[c] = n - v/2
		}
		c++
	}

	return ans

}
