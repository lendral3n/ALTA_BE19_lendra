package main

import "fmt"

func RemoveDuplicates(array []int) int {
    if len(array) == 0 {
		return 0
	}

	nextNonDuplicate := 1

	for i := 1; i < len(array); i++ {
		if array[nextNonDuplicate-1] != array[i] {
			array[nextNonDuplicate] = array[i]
			nextNonDuplicate++
		}
	}
	return nextNonDuplicate
}

func main() {
    fmt.Println(RemoveDuplicates([]int{2, 3, 3, 3, 6, 9, 9})) // 4
    fmt.Println(RemoveDuplicates([]int{2, 3, 4, 5, 6, 9, 9})) // 6
    fmt.Println(RemoveDuplicates([]int{2, 2, 2, 11})) 		  // 2
    fmt.Println(RemoveDuplicates([]int{1, 2, 3, 11, 11})) 	  // 4
}
