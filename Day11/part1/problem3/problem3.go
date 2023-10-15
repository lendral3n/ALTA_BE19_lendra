package main

import (
	"fmt"
	"sort"
)

func DragonOfLoowater(dragonHead, knightHeight []int) string {
	// your code here
	sort.Ints(dragonHead)
	sort.Ints(knightHeight)

	i := 0
	total := 0
	for _, dragon := range dragonHead {
		if i < len(knightHeight) && knightHeight[i] < dragon {
			i++
		}
		if i < len(knightHeight) && knightHeight[i] >= dragon {
			total += knightHeight[i]
			i++
		} else {
			return "Knight Fall"
		}
	}
	return fmt.Sprint(total)
}

func main() {
	fmt.Println(DragonOfLoowater([]int{5, 4}, []int{7, 8, 4}))    //11
	fmt.Println(DragonOfLoowater([]int{5, 10}, []int{5}))         // knight fall
	fmt.Println(DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2})) // knight fall
	fmt.Println(DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5})) // 10
}
