package main

import (
	"fmt"
	"reflect"
)

func isAnagram(s string, t string) bool {
	sMap := map[string]int{}
	for i := 0; i < len(s); i++ {
		sMap[string(s[i])] += 1
	}
	tMap := map[string]int{}
	for i := 0; i < len(t); i++ {
		tMap[string(t[i])] += 1
	}
	return reflect.DeepEqual(sMap, tMap)
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
