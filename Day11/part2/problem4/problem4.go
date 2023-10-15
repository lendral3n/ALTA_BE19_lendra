package main

import "fmt"

func RomanNumerals(value int) string {
	//your code here
	numerals := map[int]string{
        1: "I", 4: "IV", 5: "V", 9: "IX",
        10: "X", 40: "XL", 50: "L", 90: "XC",
        100: "C", 400: "CD", 500: "D", 900: "CM",
        1000: "M",
    }
    keys := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    result := ""
    for _, key := range keys {
        for value >= key {
            result += numerals[key]
            value -= key
        }
    }
    return result
}

func main() {
	fmt.Println(RomanNumerals(6))    //VI
	fmt.Println(RomanNumerals(9))    //IX
	fmt.Println(RomanNumerals(23))   //XXIII
	fmt.Println(RomanNumerals(2021)) //MMXXI
	fmt.Println(RomanNumerals(1646)) //MDCXLVI
}
