// Running Sum of 1d Array
// https://leetcode.com/problems/running-sum-of-1d-array/?envType=study-plan&id=level-1
package main

import "fmt"

func runningSum(nums []int) []int {
	var numsLength int = len(nums)
	var sums []int = make([]int, numsLength)

	sums[0] = nums[0]

	for i := 1; i < numsLength; i++ {
		sums[i] = nums[i] + sums[i-1]
	}
	return sums
}

func main() {
	tc := [][]int{
		{1, 2, 3, 4},
		{1, 1, 1, 1, 1},
		{3, 1, 2, 10, 1},
	}

	for _, v := range tc {
		result := runningSum(v)
		fmt.Println(result)
	}
}
