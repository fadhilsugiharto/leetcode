package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
	val := letters[0]
	for _, letter := range letters {
		if letter > target {
			return letter
		}
	}
	return val
}

func main() {
	s := "cfj"
	sb := []byte(s)
	sTarget := "a"
	target := sTarget[0]
	fmt.Println(nextGreatestLetter(sb, target))
}
