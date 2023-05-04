package main

func LevenshteinDistance(a, b string) int {
	// Write your code here.
	edits := make([][]int, len(b)+1)
	for x := range edits {
		edits[x] = make([]int, len(a)+1)
		for y := range edits[x] {
			edits[x][y] = y
		}
	}

	for i := 1; i < len(b)+1; i++ {
		edits[i][0] = edits[i-1][0] + 1
	}

	for i := 1; i < len(b)+1; i++ {
		for j := 1; j < len(a)+1; j++ {
			if b[i-1] == a[j-1] {
				edits[i][j] = edits[i-1][j-1]
				continue
			}

			edits[i][j] = 1 + min(edits[i-1][j-1], edits[i][j-1], edits[i-1][j])
		}
	}

	return edits[len(b)][len(a)]
}

func min(args ...int) int {
	cur := args[0]
	for _, n := range args {
		if cur > n {
			cur = n
		}
	}

	return cur
}
