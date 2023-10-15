package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	// your code here
	n := len(jumps)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 0; i < n; i++ {
		for j := i + 1; j <= i+2 && j < n; j++ {
			dp[j] = min(dp[j], dp[i]+abs(jumps[j]-jumps[i]))
		}
	}
	return dp[n-1]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}


func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) //40
}
