package one

import (
	"testing"
)

func TestMaxCalories(t *testing.T) {
	cases := []struct {
		name     string
		actual   string
		expected int
	}{
		{
			name:     "example",
			actual:   input,
			expected: 24000,
		},
		{
			name:     "actual input",
			actual:   input2,
			expected: 67016,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			m := Splitter(c.actual)
			_, got := MaxCalories(m)
			if c.expected != got {
				t.Errorf("expected %v, got %v", c.expected, got)
			}
		})
	}
}

func TestTopThree(t *testing.T) {
	cases := []struct {
		name     string
		actual   string
		expected int
	}{
		{
			name:     "example",
			actual:   "100 200 700",
			expected: 1000,
		},
		{
			name:     "input test",
			actual:   input,
			expected: 45000,
		},
		{
			name:     "input2 test",
			actual:   input2,
			expected: 200116,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := TopThree(c.actual)
			if c.expected != got {
				t.Errorf("expected %v, got %v", c.expected, got)
			}
		})
	}
}
