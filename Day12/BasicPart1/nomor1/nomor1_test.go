package main

import "testing"

func TestSapa(t *testing.T) {
    testCases := []struct {
        name     string
        input    string
        expected string
    }{
        {"Test 1", "Budi", "Hello Budi! Saya Golang, bahasa yang sangat menyenangkan.\n"},
        {"Test 2", "Ani", "Hello Ani! Saya Golang, bahasa yang sangat menyenangkan.\n"},
        {"Test 3", "Zaki", "Hello Zaki! Saya Golang, bahasa yang sangat menyenangkan.\n"},
        {"Test 4", "Rina", "Hello Rina! Saya Golang, bahasa yang sangat menyenangkan.\n"},
        {"Test 5", "Mira", "Hello Mira! Saya Golang, bahasa yang sangat menyenangkan.\n"},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result := Sapa(tc.input)
            if result != tc.expected {
                t.Errorf("Hasil tidak sesuai dengan yang diharapkan")
            }
        })
    }
}
