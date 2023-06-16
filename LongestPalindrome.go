package main

import "fmt"

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1

	for i < len(s)/2 {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func longestPalindrome(s string) string {
	word := string(s[0])
	for i := 0; i < len(s)-1 || len(s[i:]) > len(word); i++ {
		for j := len(s) - 1; j > i; j-- {
			if s[i] == s[j] {
				if isPalindrome(s[i : j+1]) {
					if len(s[i:j+1]) > len(word) {
						word = s[i : j+1]
					}
				}
			}
		}
	}
	return word
}

func main() {
	fmt.Println(longestPalindrome("aacabdkacaa"))
}
