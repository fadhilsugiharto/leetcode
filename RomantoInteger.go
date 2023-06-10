package main

import "fmt"

func romanToInt(s string) int {
	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	if len(s) == 1 {
		return roman[s]
	}

	num := 0
	pointer := len(s) - 1

	for pointer >= 0 {
		romanPointer := roman[string(s[pointer])]
		romanPointerLeft := roman[string(s[pointer-1])]

		if romanPointerLeft < romanPointer {
			num = num + romanPointer - romanPointerLeft
			pointer -= 2
		} else {
			num += romanPointer
			pointer -= 1
		}
		if pointer == 0 {
			return num + roman[string(s[pointer])]
		}
	}
	return num
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
