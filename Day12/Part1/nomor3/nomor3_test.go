package main

import (
	"testing"
)

func TestPrimaSegiEmpat(t *testing.T) {
    testCases := []struct {
        name     string
        wide     int
        high     int
        start    int
        expected string
    }{
        {"2 kolom, 3 baris, mulai dari 13", 2, 3, 13, "17 19\n23 29\n31 37\n150\n"},
        {"5 kolom, 2 baris, mulai dari 1", 5, 2, 1, "2 3 5 7 11\n13 17 19 23 29\n129\n"},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result := PrimaSegiEmpat(tc.wide, tc.high, tc.start)
            if result != tc.expected {
                t.Errorf("Hasil tidak sesuai dengan yang diharapkan")
            }
        })
    }
}
