package main

import (
    "fmt"
    "strings"
)

func Compare(a, b string) string {
		if strings.Contains(b, a) {
			return a
		}

		for i := len(a) - 1; i > 0; i-- {
			substr := a[0:i]
			if strings.Contains(b, substr) {
				return substr
			}
		}
	return ""
}

func main() {
    fmt.Println(Compare("AKA", "AKASHI"))     // Output: AKA
    fmt.Println(Compare("KI", "KIJANG"))      // Output: KI
    fmt.Println(Compare("KANGOORO", "KANG"))  // Output: KANG
    fmt.Println(Compare("KUPU-KUPU", "KUPU")) // Output: KUPU
    fmt.Println(Compare("ILALANG", "ILA"))    // Output: ILA
}
