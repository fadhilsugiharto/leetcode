package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	MaxCount := 1
	temp := ""
	check := true

	for i, r := range s {
		temp = string(r)
		for j := i + 1; j < len(s); j++ {
			for _, t := range temp {
				if string(s[j]) == string(t) {
					temp = ""
					check = false
					break
				}
			}
			if check == false {
				check = true
				break
			}
			temp += string(s[j])
			if MaxCount < len(temp) {
				MaxCount = len(temp)
			}
		}
	}

	return MaxCount
}

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
