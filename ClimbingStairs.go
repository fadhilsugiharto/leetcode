package main

import "fmt"

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	// list of steps needed for each staircase
	stairsStepCounts := make([]int, n+1)
	// anak tangga ne siji mung iso sak step
	stairsStepCounts[0] = 1
	// anak tangga ne 2 iso 2 step, either mletik langsung 2 langkah (langsung tekan target)
	// nek ora mletik pisan ning anak tangga 1, which then uses jumlah step yg dilakukan pas koe munggah nek anak tangga 1
	stairsStepCounts[1] = 2

	for i := 3; i <= n; i++ {
		stairsStepCounts[2] = stairsStepCounts[0] + stairsStepCounts[1]
		stairsStepCounts[0] = stairsStepCounts[1]
		stairsStepCounts[1] = stairsStepCounts[2]
	}

	return stairsStepCounts[2]
}

func main() {
	fmt.Println(climbStairs(5))
}
