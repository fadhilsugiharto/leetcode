package main

import "fmt"

func strStr(haystack string, needle string) int {
	i := 0
	hayLen := len(haystack)
	needleLen := len(needle)

	for i < len(haystack) {
		if haystack[i] == needle[0] {
			if hayLen-i >= needleLen {
				if haystack[i:i+needleLen] == needle {
					return i
				}
			} else {
				return -1
			}
		}
		i++
	}
	return -1
}

func main() {
	fmt.Println(strStr("leetcode", "leeto"))
}
