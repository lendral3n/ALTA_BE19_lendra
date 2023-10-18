package main

import "fmt"

func HitungLuasPermukaanTabung(T float64, r float64) float64 {
    const Pi = 3.14
    return 2 * Pi * r * (r + T)
}

func main() {
    // Input
    var T float64 = 20
    var r float64 = 4

    luas := HitungLuasPermukaanTabung(T, r)

    fmt.Printf("Luas Permukaan tabung: %.2f\n", luas)
}
