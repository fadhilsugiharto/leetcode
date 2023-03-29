package main

func isPalindrome(x int) bool {

	rev := 0
	ori := x

	for x > 0 {
		rev = rev*10 + x%10
		x /= 10
	}

	return rev == ori

}

//func main() {
//	fmt.Println(isPalindrome(1234321))
//}
