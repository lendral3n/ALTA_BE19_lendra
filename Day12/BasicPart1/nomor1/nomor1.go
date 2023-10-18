package main

import "fmt"

func Sapa(nama string) string {
    pesan := "Hello " + nama + "! Saya Golang, bahasa yang sangat menyenangkan.\n"
    return pesan
}

func main() {
    // Input
    var nama string
    fmt.Print("Masukkan nama: ")
    fmt.Scanln(&nama)

    fmt.Print(Sapa(nama))
}
