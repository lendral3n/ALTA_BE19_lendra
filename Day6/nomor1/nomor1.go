package main

import "fmt"

func playWithAsterik(n int) {
    for i := 1; i <= n; i++ {
        for j := 1; j <= n-i; j++ {
            fmt.Print(" ")
        }

        for j := 1; j <= i; j++ {
            fmt.Print("*")

			fmt.Print(" ")
        }

        fmt.Println()
    }
}

func main() {
    playWithAsterik(5)

	/* 
	
	* 
   * *
  * * *
 * * * *
* * * * *

	*/
}
