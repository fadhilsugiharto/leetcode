package main

import (
	"fmt"
	"strings"
	"unicode"
)

func validPalindrome(s string) bool {
	f := func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		return unicode.ToLower(r)
	}

	s = strings.Map(f, s)

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

func main() {
	fmt.Println(validPalindrome("A man, a plan, a canal: Panama"))
}
