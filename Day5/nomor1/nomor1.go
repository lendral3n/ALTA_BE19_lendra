package main

import "fmt"

func main() {
	var bilangan int
	fmt.Print("Masukan Bilangan: ")
	fmt.Scanf("%d", &bilangan)

	fmt.Println("Faktor-faktor dari bilangan", bilangan, "adalah:")
	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Println(i)
		}
	}
}