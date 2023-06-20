package main

import "fmt"

// First try
//func isPalindrome(s string) bool {
//	i, j := 0, len(s)-1
//
//	for i < len(s)/2 {
//		if s[i] != s[j] {
//			return false
//		}
//		i++
//		j--
//	}
//	return true
//}
//
//func longestPalindrome(s string) string {
//	word := string(s[0])
//	for i := 0; i < len(s)-1 || len(s[i:]) > len(word); i++ {
//		for j := len(s) - 1; j > i; j-- {
//			if s[i] == s[j] {
//				if isPalindrome(s[i : j+1]) {
//					if len(s[i:j+1]) > len(word) {
//						word = s[i : j+1]
//					}
//				}
//			}
//		}
//	}
//	return word
//}

// 5 head approach
func longestPalindrome(s string) string {
	lenS := len(s)
	if lenS == 0 {
		return ""
	}

	var l, r, pl, pr int

	for r < lenS {
		// For even string length
		for r+1 < lenS && s[l] == s[r+1] {
			r++
		}

		//	For odd string length
		for l-1 >= 0 && r+1 < lenS && s[l-1] == s[r+1] {
			l--
			r++
		}

		if r-l > pr-pl {
			pl, pr = l, r
		}

		l = (l+r)/2 + 1
		r = l
	}
	return s[pl : pr+1]
}

func main() {
	fmt.Println(longestPalindrome("aacabdkacaa"))
}
