package main

import (
	"testing"
)
func TestIsPrime(t *testing.T) {
    if !IsPrime(2) {
		t.Errorf("ApakahPrima(2) = false; seharusnya true")
	}
	if IsPrime(4) {
		t.Errorf("ApakahPrima(4) = true; seharusnya false")
	}
}


func TestPrimeX(t *testing.T) {
	if PrimeX(9) != 23 {
		t.Errorf("PrimeX(9) = %d; seharusnya 23", PrimeX(9))
	}
	if PrimeX(10) != 29 {
		t.Errorf("PrimeX(10) = %d; seharusnya 29", PrimeX(10))
	}
}
