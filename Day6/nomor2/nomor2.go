package main

import (
	"fmt"
)

func DrawXYZ(n int) string {
	result :=  ""
	for i := 1; i <= n*n; i++ {
		if i%3 == 0 {
			result += "X"
		} else if i%2 == 0 {
			result += "Z"
		} else {
			result += "Y"
		}

		if i%n == 0 {
			result += "\n"
		} else {
			result += " "
		}
	}
	return result
}


func main() {

	fmt.Println(DrawXYZ(3))

	/*
	Y Z X
	Z Y X
	Y Z X
	*/

	fmt.Println(DrawXYZ(5))

	/*
	Y Z X Z Y
	X Y Z X Z
	Y X Y Z X
	Z Y X Y Z
	X Z Y X Y
	*/
}