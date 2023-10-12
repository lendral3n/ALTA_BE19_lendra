package main

import (
	"fmt"
	"strconv"
)

func isPrime(n int) bool {
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

func primaSegiEmpat(wide, high, start int) string {
	primes := []int{}
	for i := start + 1; len(primes) < wide*high; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	result := ""
	sum := 0
	for i := 0; i < high; i++ {
		for j := 0; j < wide; j++ {
			index := i*wide + j
			result += strconv.Itoa(primes[index]) + " "
			sum += primes[index]
		}
		result += "\n"
	}
	result += strconv.Itoa(sum)
	return result
}

func main() {
	fmt.Println(primaSegiEmpat(2, 3, 13))
	/*
        17 19
        23
        29
        31 37
    */
	fmt.Println(primaSegiEmpat(5, 2, 1))
	/*
        2 3 5 7 11
        13 17 19 23 29
    */
}
