package aoc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Parsing(t *testing.T) {
	t.Skipf("Skipping until code is ready")
	cases := []struct {
		name     string
		actual   string
		expected int
	}{
		{
			name:     "Part1 input",
			actual:   "part1.input",
			expected: 0,
		},
		{
			name:     "Part2 input",
			actual:   "part2.input",
			expected: 0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			in, err := readInput(c.actual)
			assert.NoError(t, err)
			got := parseInput(in)
			assert.Equal(t, c.expected, got)
		})
	}
}

func Test_Inputs(t *testing.T) {
	t.Skipf("Skipping until code is ready")
	cases := []struct {
		name     string
		actual   string
		expected int
	}{
		{
			name:     "Part1 input",
			actual:   "part1.input",
			expected: 0,
		},
		{
			name:     "Part2 input",
			actual:   "part2.input",
			expected: 0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			in, err := readInput(c.actual)
			fmt.Println("input is:", string(in))
			assert.NoError(t, err)
		})
	}
}
