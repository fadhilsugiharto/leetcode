package main

import "fmt"

func roman(s string) int {
	num := 0

	if string(s) == "M" {
		num += 1000
	} else if string(s) == "D" {
		num += 500
	} else if string(s) == "C" {
		num += 100
	} else if string(s) == "L" {
		num += 50
	} else if string(s) == "X" {
		num += 10
	} else if string(s) == "V" {
		num += 5
	} else if string(s) == "I" {
		num += 1
	}

	return num
}

func romanToInt(s string) int {
	if len(s) == 1 {
		return roman(s)
	}

	num := 0
	pointer := len(s) - 1

	for pointer >= 0 {
		romanPointer := roman(string(s[pointer]))
		romanPointerLeft := roman(string(s[pointer-1]))

		if romanPointerLeft < romanPointer {
			num = num + romanPointer - romanPointerLeft
			pointer -= 2
		} else {
			num += romanPointer
			pointer -= 1
		}
		if pointer == 0 {
			return num + roman(string(s[pointer]))
		}
	}

	return num
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
