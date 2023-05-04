package main

import (
	"fmt"
	"strings"
)

var (
	lessThan20 = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven",
		"Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen",
		"Sixteen", "Seventeen", "Eighteen", "Nineteen"}

	tens = []string{"", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}

	thousands = []string{"", "Thousand", "Million", "Billion"}
)

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	words, i := "", 0
	for num > 0 {
		mod := num % 1000
		if mod != 0 {
			words = getNumberString(mod) + thousands[i] + " " + words
		}

		num = num / 1000
		i++
	}

	return strings.TrimSpace(words)
}

func getNumberString(num int) string {
	if num == 0 {
		return ""
	}

	if num < 20 {
		return lessThan20[num] + " "
	}

	if num < 100 {
		return tens[num/10] + " " + getNumberString(num%10)
	}

	return lessThan20[num/100] + " Hundred " + getNumberString(num%100)
}

func main() {
	nums := []int{2156056, 123, 12456}
	for _, num := range nums {
		word := numberToWords(num)
		fmt.Println(fmt.Sprintf("Input: %d\nOutput: %s", num, word))
		fmt.Println()
	}
}
