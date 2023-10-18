package main

import (
	"testing"
)

func TestHitungLuasPermukaanTabung(t *testing.T) {
    testCases := []struct {
        T, r, expected float64
    }{
        {20, 4, 301.44},
        {10, 10, 628.0},
        {15, 30, 2826.0},
        {5, 5, 157.0},
        {100, 200, 125600.0},
    }

    for _, tc := range testCases {
        result := HitungLuasPermukaanTabung(tc.T, tc.r) // Mengubah "HitungLuasPermukaanTabung" menjadi "hitungLuasPermukaanTabung"
        if result != tc.expected {
            t.Error("Test case failed for T: ", tc.T, ", r: ", tc.r,
                     ". Expected: ", tc.expected, ", but got: ", result)
        }
    }
}
