package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	prefix := ""
	first := strs[0]

	for i := 0; i < len(first); i++ {
		checker := ""
		for i2, str := range strs {
			if i2 == 0 {
				checker = string(str[i])
			} else {
				if len(str) == i || string(str[i]) != checker {
					return prefix
				}
			}
		}
		prefix += string(first[i])
	}
	return prefix
}

func main() {
	strs := []string{
		"ab", "a",
	}
	fmt.Println(longestCommonPrefix(strs))
}
