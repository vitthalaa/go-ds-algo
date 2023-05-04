package main

import "fmt"

var (
	mapping = map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
)

func letterCombinations(digits string) []string {
	var res []string
	for i := 0; i < len(digits); i++ {
		mp, ok := mapping[string(digits[i])]
		if !ok || len(mp) == 0 {
			continue
		}

		if len(res) == 0 {
			// can do res = m directly, but it seems modifies mapping after modifying res
			for _, m := range mp {
				res = append(res, m)
			}

			continue
		}

		for i, s := range res {
			for j, m := range mp {
				if j == 0 {
					res[i] = s + m
					continue
				}

				res = append(res, s+m)
			}
		}
	}

	return res
}

func main() {
	res := letterCombinations("23")
	fmt.Println(res)
}
