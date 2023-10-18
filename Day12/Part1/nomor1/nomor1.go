package main

import "fmt"

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func PrimeX(number int) int {
	count := 0
	for i := 2; ; i++ {
		if IsPrime(i) {
			count++
			if count == number {
				return i
			}
		}
	}
}

func main() {
	fmt.Println(PrimeX(1))  // 2
	fmt.Println(PrimeX(5))  // 11
	fmt.Println(PrimeX(8))  // 19
	fmt.Println(PrimeX(9))  // 23
	fmt.Println(PrimeX(10)) // 29
}