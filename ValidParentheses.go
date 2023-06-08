package main

import "fmt"

func isValid(s string) bool {
	if len(s)%2 != 0 || len(s) == 0 {
		return false
	}

	pairs := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}

	var stack []rune

	for _, i := range s {
		if _, exists := pairs[i]; exists {
			stack = append(stack, i)
		} else if len(stack) == 0 || pairs[stack[len(stack)-1]] != i {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("(){}{}"))
}
