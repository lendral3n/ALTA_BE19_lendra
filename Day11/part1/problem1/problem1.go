package main

import (
	"fmt"
)

func SimpleEquations(a, b, c int) interface{} {
	// your code here
	for x := 1; x <= 100; x++ {
		for y := 1; y <= 100; y++ {
				for z := 1; z <= 100; z++ {
					if x+y+z == a && x*y*z == b && x*x+y*y+z*z == c {
						return []int{x, y, z}
					}
				}
			}
		}
		return "no solution"
	}


func main() {
	fmt.Println(SimpleEquations(1, 2, 3))  // no solution
	fmt.Println(SimpleEquations(6, 6, 14)) // [1,2,3]
}
