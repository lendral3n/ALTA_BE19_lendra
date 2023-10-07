package main

import "fmt"

func findMaxSumSubArray(k int, arr []int) int {
    maxSum := 0
	windowSum := 0
	windowStart := 0

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		windowSum += arr[windowEnd]

		if windowEnd >= k-1 {
			if windowSum > maxSum {
				maxSum = windowSum
			}

			windowSum -= arr[windowStart]
			windowStart++
		}
	}
	return maxSum
}

func main() {
    fmt.Println(findMaxSumSubArray(3, []int{2, 1, 5, 1, 3, 2})) // Output: 9
    fmt.Println(findMaxSumSubArray(2, []int{2, 3, 4, 1, 5}))    // Output: 7
    fmt.Println(findMaxSumSubArray(2, []int{2, 1, 4, 1, 1}))    // Output: 5
    fmt.Println(findMaxSumSubArray(3, []int{2, 1, 4, 1, 1}))    // Output: 7
    fmt.Println(findMaxSumSubArray(4, []int{2, 1, 4, 1, 1}))    // Output: 8
}
