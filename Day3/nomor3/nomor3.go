package main

import "fmt"

func main() {
    // Input
    var T float64 = 20
	var r float64 = 4

	const Pi = 3.14

	luas := 2 * Pi * r * (r + T)

	fmt.Printf("Luas Permukaan tabung: %.2f\n", luas)
}
