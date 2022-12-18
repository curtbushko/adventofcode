package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Run(t *testing.T) {
	cases := []struct {
		name     string
		actual   string
		expected int
	}{
		{
			name:     "example",
			actual:   "example.input",
			expected: 13,
		},
		{
			name:     "large example",
			actual:   "large.input",
			expected: 6175,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := run(c.actual)
			assert.Equal(t, c.expected, got)
		})
	}
}
