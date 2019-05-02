package main

import (
	"math"
	"fmt"
)
// https://leetcode.com/problems/minimum-path-sum

func main() {
	
	grid := [][]int {
		{1,2,5},
  		{3,2,1},
	}

	sum := minPathSum(grid)

	fmt.Println(sum)
}

func minPathSum(grid [][]int) int {
    return minPathSumPos(grid, 0, 0)
}

func minPathSumPos(grid [][]int, i int, j int) int {
    
    if i >= len(grid[0]) || j >= len(grid)  {
        return math.MaxInt8
    } 
    
    if i == len(grid[0])  - 1 && j == len(grid) - 1 {
        return grid[j][i]
	} 

    return grid[j][i] + min(minPathSumPos(grid, i + 1, j), minPathSumPos(grid, i, j + 1)) 
}

func min(x, y int) int {
    if x > y {
        return y
    }
    return x
}