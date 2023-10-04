package main

import "fmt"

func main() {
	var bilangan int
	fmt.Print("Masukan Bilangan: ")
	fmt.Scanf("%d", &bilangan)

	fmt.Println("Faktor-faktor dari bilangan", bilangan, "adalah:")
		faktor := []int{}

	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			faktor = append(faktor, i)
		}
	}

	for i := len(faktor) - 1; i >= 0; i-- {
		fmt.Println(faktor[i])
	}
}