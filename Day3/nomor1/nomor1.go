package main

import "fmt"

func main() {
    // Input
    var nama string
    fmt.Print("Masukkan nama: ")
    fmt.Scanln(&nama)

    fmt.Printf("Hello %s! Saya Golang, bahasa yang sangat menyenangkan.\n", nama)
}
