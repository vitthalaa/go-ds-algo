package main

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	curMax := nums[0]
	for i := 1; i < len(nums); i++ {
		curMax = max(curMax+nums[i], nums[i])
		maxSum = max(maxSum, curMax)
	}

	return maxSum
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	// The subarray [4,-1,2,1] has the largest sum 6.
	ms := maxSubArray(nums)

	fmt.Println(ms)
}
