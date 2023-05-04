package main

import "math"

func MinNumberOfCoinsForChange(n int, denoms []int) int {
	// Write your code here.
	if n == 0 {
		return 0
	}

	// n = 7, denoms = [1, 5, 10], op = 3
	// coins = [0, max, max, max, max, max, max, max]

	// denom = 1
	// for amount 0 to 7
	// 1[_, 1, _, _, _, _, _, _]
	// 2[_, _, 2, _, _, _ _, _]
	// 3[_, _, _, 3, _, _, _, _]
	// 4[_, _, _, _, 4, _, _, _]
	// 5[_, _, _, _, _, 5, _, _]
	// 6[_, _, _, _, _, _, 6, _]
	// 7[_, _, _, _, _, _, _, 7]
	// coins = [0, 1, 2, 3, 4, 5, 6, 7]

	// denom = 5
	// 5[_, _, _, _, _, 1, _, _] => min(coins[5], coins[5-5]+1)
	// 6[_, _, _, _, _, _, 2, _] => min(coins[6], coins[6-5]+1)
	// 7[_, _, _, _, _, _, _, 3] => min(coins[7], coins[7-5]+1)
	// coins = [0, 1, 2, 3, 4, 1, 2, 3]

	// so formula is coins[amount] = min(coins[amount], coins[amount - denom] + 1)

	// Initialize an array of the minimum number coins to make change for all amounts between 0 and n inclusive.
	coins := make([]int, n+1)
	// keep 0th index value to 0, remaining ones to maximum
	for i := 1; i < len(coins); i++ {
		coins[i] = math.MaxInt32
	}

	for i := 0; i < len(denoms); i++ {
		for amount := 0; amount <= n; amount++ {
			if amount < denoms[i] {
				continue
			}

			coins[amount] = min(coins[amount], coins[amount-denoms[i]]+1)
		}
	}

	if coins[n] == math.MaxInt32 {
		return -1
	}

	return coins[n]
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}
