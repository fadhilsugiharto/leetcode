package main

import "fmt"

//func lengthOfLongestSubstring(s string) int {
//	if s == "" {
//		return 0
//	}
//	MaxCount := 1
//	temp := ""
//	check := true
//
//	for i, r := range s {
//		temp = string(r)
//		for j := i + 1; j < len(s); j++ {
//			for _, t := range temp {
//				if string(s[j]) == string(t) {
//					temp = ""
//					check = false
//					break
//				}
//			}
//			if check == false {
//				check = true
//				break
//			}
//			temp += string(s[j])
//			if MaxCount < len(temp) {
//				MaxCount = len(temp)
//			}
//		}
//	}
//
//	return MaxCount
//}

func lengthOfLongestSubstring(s string) int {
	dict := [128]bool{}
	count := 0
	max := 0

	for i, j := 0, 0; i < len(s); i++ {
		if dict[s[i]] == true {
			for ; dict[s[i]]; j++ {
				count--
				dict[s[j]] = false
			}
		}
		dict[s[i]] = true
		count++
		if max < count {
			max = count
		}
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("bbbbbbb"))
}
