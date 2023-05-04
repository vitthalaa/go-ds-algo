package main

func MaxSubsetSumNoAdjacent(array []int) int {
	// Write your code here.
	if len(array) == 0 {
		return 0
	}

	if len(array) == 1 {
		return array[0]
	}

	// create array of max sums at every index of input array
	// max sum at position i is max of (max sum of previous element and max sum of previous non-adjacent element + current element)
	maxSums := make([]int, len(array))

	// max sum at first position is element itself
	maxSums[0] = array[0]
	// max sum at second position is second position is max of second element and first element
	maxSums[1] = max(array[0], array[1])

	for i := 2; i < len(array); i++ {
		maxSums[i] = max(maxSums[i-1], maxSums[i-2]+array[i])
	}

	return maxSums[len(maxSums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
