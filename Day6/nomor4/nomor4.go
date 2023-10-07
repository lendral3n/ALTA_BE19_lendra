package main

import "fmt"

func ubahHuruf(sentence string) string {
    alfabetMap := make(map[rune]rune)

	alfabetAsli := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alfabetGeser := "KLMNOPQRSTUVWXYZABCDEFGHIJ"

	for i, karakter := range alfabetAsli {
		alfabetMap[karakter] = rune(alfabetGeser[i])
	}

	kalimatTerenkripsi := ""

	for _, karakter := range sentence {
		if karakter == ' ' {
			kalimatTerenkripsi += " "
		} else {
			kalimatTerenkripsi += string(alfabetMap[karakter])
		}
	}
	return kalimatTerenkripsi
}

func main() {
    fmt.Println(ubahHuruf("SEPULSA OKE"))
    			// Output: COZEVCK YUO
    fmt.Println(ubahHuruf("ALTERRA ACADEMY"))
    			// Output: KVDOBBK KMKNOWI
    fmt.Println(ubahHuruf("INDONESIA"))
    			// Output: SXNYXOCSK
    fmt.Println(ubahHuruf("GOLANG"))
    			// Output: QYVKXQ
    fmt.Println(ubahHuruf("PROGRAMMER"))
    			// Output: ZBYQBKWWOB
}
