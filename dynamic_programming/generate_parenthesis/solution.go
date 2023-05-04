package main

import "fmt"

func GenerateParenthesis(n int) []string {
	res := []string{}
	str := make([]byte, 0, n*2)
	dfs(&res, str, 0, 0, n)

	return res
}

func dfs(res *[]string, str []byte, open, close, max int) {
	if open == max && close == max {
		*res = append(*res, string(str))
		return
	}

	if open < max {
		dfs(res, append(str, '('), open+1, close, max)
	}

	if close < open {
		dfs(res, append(str, ')'), open, close+1, max)
	}
}

func main() {
	n := 3
	res := GenerateParenthesis(n)

	fmt.Println(fmt.Sprintf("Input: %d\nOutput: %v", n, res))
}
