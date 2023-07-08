package main

import "fmt"

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if s == goal {
		letterCount := map[int32]int{}
		for _, letter := range s {
			letterCount[letter] += 1
			if letterCount[letter] == 2 {
				return true
			}
		}
	} else {
		diffIndex := []int{}

		for i, _ := range s {
			if s[i] != goal[i] {
				diffIndex = append(diffIndex, i)
				if len(diffIndex) == 3 {
					return false
				}
			}
		}

		if len(diffIndex) == 2 && s[diffIndex[0]] == goal[diffIndex[1]] && s[diffIndex[1]] == goal[diffIndex[0]] {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(buddyStrings("ab", "babbb"))
}
