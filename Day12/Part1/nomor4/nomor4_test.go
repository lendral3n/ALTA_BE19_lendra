package main

import "testing"

func TestMaxSequence(t *testing.T) {
    testCases := []struct {
        name     string
        input    []int
        expected int
    }{
        {"Contoh 1", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
        {"Contoh 2", []int{-2, -5, 6, -2, -3, 1, 5, -6}, 7},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result := MaxSequence(tc.input)
            if result != tc.expected {
                t.Errorf("Hasil tidak sesuai dengan yang diharapkan")
            }
        })
    }
}
