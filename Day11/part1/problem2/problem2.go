package main

import "fmt"

func MoneyChange(money int) []int {
	// your code here
	var result []int
	change := []int{10000, 5000, 2000, 1000, 500, 200, 100, 50, 20, 10, 5, 1}

	for _, change := range change {
		for money >= change {
			money -= change
			result = append(result, change)
		}
	}
	return result
}

func main() {
	fmt.Println(MoneyChange(123))   //[100 20 1 1 1]
	fmt.Println(MoneyChange(432))   //[200 200 20 10 1 1]
	fmt.Println(MoneyChange(543))   //[500 20 20 1 1 1]
	fmt.Println(MoneyChange(7752))  //[5000 2000 500 200 50 1 1]
	fmt.Println(MoneyChange(15321)) //[10000 5000 200 100 20 1]
}
