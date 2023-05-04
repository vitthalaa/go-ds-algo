package main

import (
	"fmt"
)

func sqrtInt(x int) int {
	if x == 0 {
		return 0
	}

	start := 1
	end := x
	ans := 0

	for start <= end {
		mid := start + (end-start)/2
		if mid <= x/mid {
			start = mid + 1
			ans = mid
			continue
		}

		end = mid - 1
	}

	return ans
}

func sqrtFloat(x float64) float64 {
	start := float64(1)
	end := x
	precise := 1e-12

	if end < 1 {
		end = 1
	}

	for end-start >= precise {
		mid := start + (end-start)/2
		if mid*mid < x {
			start = mid
			continue
		}

		end = mid
	}

	return start
}

func main() {
	s := sqrtInt(9)
	fmt.Println(s)

	f := sqrtFloat(2)
	fmt.Println(f)
}
