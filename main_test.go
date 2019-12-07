package main

import (
	"testing"
)

func TestMassCalculation(t *testing.T) {

	var testData = []struct {
		mass     float64
		expected int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, tt := range testData {
		result := calculateFuel(tt.mass)
		if result != tt.mass {
			t.Errorf("Result: %v, should be %v\n", result, tt.expected)
		}
	}
}
