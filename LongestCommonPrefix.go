package main

import "fmt"

//func longestCommonPrefix(strs []string) string {
//	prefix := ""
//	first := strs[0]
//
//	for i := 0; i < len(first); i++ {
//		checker := ""
//		for i2, str := range strs {
//			if i2 == 0 {
//				checker = string(str[i])
//			} else {
//				if len(str) == i || string(str[i]) != checker {
//					return prefix
//				}
//			}
//		}
//		prefix += string(first[i])
//	}
//	return prefix
//}

func longestCommonPrefix(strs []string) string {
	p := strs[0]
	for _, str := range strs {
		i := 0
		if str == p {
			continue
		}
		for ; len(str) > i && len(p) > i && str[i] == p[i]; i++ {
		}
		p = p[:i]
	}
	return p
}

func main() {
	strs := []string{
		"a", "abcd",
	}
	fmt.Println(longestCommonPrefix(strs))
}
