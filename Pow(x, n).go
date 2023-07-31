package main

import (
	"fmt"
)

//func myPow(x float64, n int) float64 {
//	if n > 0 {
//		ans := x
//		for i := 0; i < n-1; i++ {
//			ans = ans * x
//		}
//		return ans
//	} else if n < 0 {
//		ans := x
//		for i := 0; i > n+1; i-- {
//			ans = ans * x
//		}
//		ans = 1 / ans
//		return ans
//	} else {
//		return 1.0
//	}
//}

func power(x float64, n int) float64 {
	if x == 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	ans := power(x, n/2)
	ans = ans * ans

	if n%2 != 0 {
		ans *= x
	}

	return ans
}

func myPow(x float64, n int) float64 {
	if n < 0 {
		n *= -1
		return 1 / power(x, n)
	}
	return power(x, n)
}

func main() {
	fmt.Print(myPow(2, 5))
}
