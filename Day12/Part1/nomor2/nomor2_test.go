package main

import ("testing"
		"strconv")

func TestFibonacci(t *testing.T) {
    testCases := []struct {
        name     string
        input    int
        expected int
    }{
        {"Fibonacci ke-0", 0, 0},
        {"Fibonacci ke-2", 2, 1},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result := Fibonacci(tc.input)
            if result != tc.expected {
                t.Error("fibonacci dari " + tc.name + " = " + strconv.Itoa(result) + "; seharusnya " + strconv.Itoa(tc.expected))
            }
        })
    }
}
