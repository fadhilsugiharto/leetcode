package main

import (
	"fmt"
)

// WITH MAP
//func isAnagram(s string, t string) bool {
//	sMap := map[string]int{}
//	for i := 0; i < len(s); i++ {
//		sMap[string(s[i])] += 1
//	}
//	tMap := map[string]int{}
//	for i := 0; i < len(t); i++ {
//		tMap[string(t[i])] += 1
//	}
//	return reflect.DeepEqual(sMap, tMap)
//}

// Sort and compare
//func isAnagram(s string, t string) bool {
//	// Convert the string to a slice of runes
//	runesS := []rune(s)
//
//	// Sort the rune slice
//	sort.Slice(runesS, func(i, j int) bool {
//		return runesS[i] < runesS[j]
//	})
//
//	// Convert the string to a slice of runes
//	runesT := []rune(t)
//
//	// Sort the rune slice
//	sort.Slice(runesT, func(i, j int) bool {
//		return runesT[i] < runesT[j]
//	})
//
//	return reflect.DeepEqual(runesS, runesT)
//}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false

	}
	var ans [26]int
	for i := 0; i < len(s); i++ {
		ans[s[i]-'a']++
		ans[t[i]-'a']--
	}

	for _, v := range ans {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
