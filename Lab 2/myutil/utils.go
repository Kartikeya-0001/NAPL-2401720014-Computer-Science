package myutil

import "strings"

func Reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func CountVowels(s string) int {
	count := 0

	for _, ch := range strings.ToLower(s) {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			count++
		}
	}

	return count
}

func Factorial(n int) int {
	result := 1

	for i := 1; i <= n; i++ {
		result *= i
	}

	return result
}

func Power(a, b int) int {
	result := 1

	for i := 0; i < b; i++ {
		result *= a
	}

	return result
}
