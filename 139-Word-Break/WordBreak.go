package main

import "fmt"

//func wordBreak(s string, wordDict []string) bool {
//	i := 0
//	strLen := len(s)
//
//	for i < len(s) {
//		found := false
//		for _, word := range wordDict {
//			n := len(word)
//			if strLen-i < n {
//				continue
//			}
//			if s[i:i+n] == word {
//				found = true
//				i += n
//				break
//			}
//		}
//		if found == false {
//			return false
//		}
//	}
//	return true
//}

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[n] = true

	for i := n - 1; i >= 0; i-- {
		for _, word := range wordDict {
			wordLen := len(word)
			if n-i >= wordLen && s[i:i+wordLen] == word && dp[i+wordLen] == true {
				dp[i] = true
				break
			}
		}
	}
	return dp[0]
}

func main() {
	fmt.Println(wordBreak("cars", []string{"car", "ca", "rs"}))
}
