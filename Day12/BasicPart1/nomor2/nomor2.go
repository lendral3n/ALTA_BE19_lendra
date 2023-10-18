package main

import "fmt"

func HitungLuasSegitiga(alas float64, tinggi float64) float64 {
    return 0.5 * alas * tinggi
}

func main() {
    // Input
    var alas float64 = 20
    var tinggi float64 = 25

    luas := HitungLuasSegitiga(alas, tinggi)

    fmt.Println("Luas Segitiga:", luas)
}
