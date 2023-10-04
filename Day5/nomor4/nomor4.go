package main

import "fmt"

func palindrome(input string) bool {
	for i, j := 0, len(input)-1; i < j; {
		if input[i] != input[j] {
			return false
		}
		i++
		j--
	}
	return true
}


func main() {
	fmt.Println(palindrome("civic"))	// true
	fmt.Println(palindrome("katak"))	// true
	fmt.Println(palindrome("kasur rusak"))	// true
	fmt.Println(palindrome("kupu-kupu"))	// false
	fmt.Println(palindrome("lion"))	// false
}