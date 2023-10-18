package main

import (
	"reflect"
	"testing"
)


func TestPlayingDomino(t *testing.T) {
    testCases := []struct {
        cards [][]int
        deck []int
        expected []int
    }{
        {[][]int{{6, 5}, {3, 4}, {2, 1}, {3, 3}}, []int{4, 3}, []int{3, 4}},
        {[][]int{{6, 5}, {3, 3}, {3, 4}, {2, 1}}, []int{3, 6}, []int{6, 5}},
    }

    for _, tc := range testCases {
        result := PlayingDomino(tc.cards, tc.deck)
        if !reflect.DeepEqual(result, tc.expected) {
            t.Error("Test case failed for cards: ", tc.cards,
                     ", deck: ", tc.deck,
                     ". Expected: ", tc.expected,
                     ", but got: ", result)
        }
    }
}
