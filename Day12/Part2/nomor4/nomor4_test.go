package main

import 
	"testing"

func TestMostAppearItem(t *testing.T) {
    testCases := []struct {
        items []string
        expected string
    }{
        {[]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}, "golang->1 ruby->2 js->4 "},
        {[]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}, "C->1 D->2 B->3 A->4 "},
    }

    for _, tc := range testCases {
        result := MostAppearItem(tc.items)
        if result != tc.expected {
            t.Error("Test case failed for items: ", tc.items,
                     ". Expected: ", tc.expected,
                     ", but got: ", result)
        }
    }
}
