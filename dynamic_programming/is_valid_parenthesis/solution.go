package main

import "fmt"

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}

	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := []rune{}

	for _, r := range s {
		// if open then append to stack
		if _, ok := pairs[r]; ok {
			stack = append(stack, r)
			continue
		}

		// if close and stack is empty means there is no open pair
		// or if close one does not match with pair
		if len(stack) == 0 || pairs[stack[len(stack)-1]] != r {
			return false
		}

		// remove last element
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("(())"))
}
