package main

import (
	"testing"
)

func Test_CreateGrid(t *testing.T) {
	cases := []struct {
		name     string
		actual   string
		expected int
	}{
		{
			name:     "example",
			actual:   "example.input",
			expected: 31,
		},
		// {
		// 	name:     "large example",
		// 	actual:   "large.input",
		// 	expected: 6175,
		// },
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := createGraph(c.actual)
			got.Print()
		})
	}
}
