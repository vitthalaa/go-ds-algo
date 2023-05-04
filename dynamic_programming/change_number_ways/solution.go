package main

func NumberOfWaysToMakeChange(n int, denoms []int) int {
	// Write your code here.
	if n == 0 {
		return 1
	}

	// n = 6, denoms = [1, 5], op = 2 => 1x1 + 1x5 and 6x1

	// Initialize an array of the number of ways to make change for all amounts between 0 and n inclusive.
	// Note that there is only one way to make change for 0: that is to not use any coins.
	// ways = [1, 0, 0, 0, 0, 0, 0]

	// denom = 1
	// for amount 0 to 6
	// 1[_, 1, _, _, _, _, _]
	// 2[_, _, 1, _, _, _ _]
	// 3[_, _, _, 1, _, _, _]
	// 4[_, _, _, _, 1, _, _]
	// 5[_, _, _, _, _, 1, _]
	// 6[_, _, _, _, _, _, 1]
	// ways = [1, 1, 1, 1, 1, 1, 1]

	// denom = 5
	// 5[_, _, _, _, _, 2, _] => ways[5] + ways[5-5]
	// 6[_, _, _, _, _, _, 2] => ways[6] + ways[6-5]
	// ways = [1, 1, 1, 1, 1, 2, 2]

	// Initialize an array of the number of ways to make change for all amounts between 0 and n inclusive.
	ways := make([]int, n+1)
	// Note that there is only one way to make change for 0: that is to not use any coins.
	ways[0] = 1

	for i := 0; i < len(denoms); i++ {
		for amount := 0; amount <= n; amount++ {
			if amount < denoms[i] {
				continue
			}

			rem := amount - denoms[i]
			ways[amount] = ways[amount] + ways[rem]
		}
	}

	return ways[n]
}
