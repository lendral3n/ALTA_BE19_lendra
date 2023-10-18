package main

import "testing"

func TestMaximumBuyProduct(t *testing.T) {
	tests := []struct {
		money    int
		prices   []int
		expected int
	}{
		{50000, []int{25000, 25000, 10000, 14000}, 3},
		{30000, []int{15000, 10000, 12000, 5000, 3000}, 4},
		{10000, []int{2000, 3000, 1000, 2000, 10000}, 4},
		{4000, []int{7500, 3000, 2500, 2000}, 1},
		{0, []int{10000, 30000}, 0},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := MaximumBuyProduct(test.money, test.prices)
			if result != test.expected {
				t.Errorf("Expected %d, but got %d", test.expected, result)
			}
		})
	}
}
