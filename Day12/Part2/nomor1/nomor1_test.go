package main

import 
    "testing"

func TestFindMinAndMax(t *testing.T) {
    testCases := []struct {
        arr []int
        expected string
    }{
        {[]int{5, 7, 4, -2, -1, 8}, "min: -2 index: 3 max: 8 index: 5"},
        {[]int{2, -5, -4, 22, 7, 7}, "min: -5 index: 1 max: 22 index: 3"},
    }

    for _, tc := range testCases {
        result := FindMinAndMax(tc.arr)
        if result != tc.expected {
            t.Error("Test case failed for arr: ", tc.arr,
                     ". Expected: ", tc.expected, ", but got: ", result)
        }
    }
}
