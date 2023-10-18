package main

import "testing"

func TestHitungLuasSegitiga(t *testing.T) {
    testCases := []struct {
        alas, tinggi, expected float64
    }{
        {20, 25, 250},
        {10, 10, 50},
        {15, 30, 225},
        {5, 5, 12.5},
        {100, 200, 10000},
    }

    for _, tc := range testCases {
        result := HitungLuasSegitiga(tc.alas, tc.tinggi)
        if result != tc.expected {
            t.Error("Test case failed for alas: ", tc.alas, ", tinggi: ", tc.tinggi,
                     ". Expected: ", tc.expected, ", but got: ", result)
        }
    }
}
