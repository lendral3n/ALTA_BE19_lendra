package main

import (
	"sort"
)

func MaximumBuyProduct(money int, productPrice []int) int {
	sort.Ints(productPrice)
	count := 0
	for _, price := range productPrice {
		if money >= price {
			money -= price
			count++
		} else {
			break
		}
	}
	return count
}
