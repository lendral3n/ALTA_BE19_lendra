package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Print("Hallo")

	var a int
	fmt.Scanln(&a)
	fmt.Println("Hasil input a -", a)

	var b string
	fmt.Scanln(&b)
	fmt.Println("Hasil input b -", b)

	var nama string = "Lendra"
	var level int = 3
	var umur int = 18

	fmt.Println(nama,level,umur)
}