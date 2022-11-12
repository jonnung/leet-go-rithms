// Find Pivot Index
// https://leetcode.com/problems/find-pivot-index/?envType=study-plan&id=level-1
package main

import "fmt"

func sliceSum(nums []int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	return sum
}

func pivotIndex(nums []int) int {
	var numsLength int = len(nums)
	for i := 0; i < numsLength; i++ {
		leftSum := 0
		rightSum := 0

		if i > 0 {
			leftSum = sliceSum(nums[:i])
		}
		rightSum = sliceSum(nums[i+1 : numsLength])

		if leftSum == rightSum {
			return i
		}
	}
	return -1
}

func main() {
	tc := [][]int{
		{1, 7, 3, 6, 5, 6},
		{1, 2, 3},
		{2, 1, -1},
	}

	for _, v := range tc {
		result := pivotIndex(v)
		fmt.Println(result)
	}
}
