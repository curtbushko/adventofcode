package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func Test_Run(t *testing.T) {
// 	cases := []struct {
// 		name     string
// 		actual   string
// 		expected int
// 	}{
// 		{
// 			name:     "example",
// 			actual:   "example.input",
// 			expected: 0,
// 		},
// 		// {
// 		// 	name:     "large example",
// 		// 	actual:   "large.input",
// 		// 	expected: 6175,
// 		// },
// 	}
// 	for _, c := range cases {
// 		t.Run(c.name, func(t *testing.T) {
// 			got := run(c.actual)
// 			assert.Equal(t, c.expected, got)
// 		})
// 	}
// }

func Test_MakeOperations(t *testing.T) {
	cases := []struct {
		name     string
		actual   string
		expected []int
	}{
		{
			name:     "small ops",
			actual:   smallops,
			expected: []int{0, 0, 0, 2, 0, 0, 3, 0},
		},
		{
			name:     "medium ops",
			actual:   mediumops,
			expected: []int{0, 15, 0, -11, 0, 6, 0, -3, 0, 5, 0, -1, 0, -8, 0, 13, 0, 4, 0, 0, -1, 0, 5, 0, -1},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := makeOperations(c.actual)
			assert.Equal(t, c.expected, got)
		})
	}
}

func Test_Operations(t *testing.T) {
	cases := []struct {
		name     string
		actual   []int
		expected int
	}{
		{
			name:     "example",
			actual:   []int{0, 15, 0, -11, 0, 6, 0, -3, 0, 5, 0, -1, 0, -8, 0, 13, 0, 4, 0, 0, -1, 0, 5, 0, -1},
			expected: 420,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := processOperations(c.actual)
			assert.Equal(t, c.expected, got)
		})
	}
}

func Test_Run(t *testing.T) {
	cases := []struct {
		name     string
		actual   string
		expected int
	}{
		{
			name:     "example input",
			actual:   "example.input",
			expected: 13140,
		},
		{
			name:     "large input",
			actual:   "large.input",
			expected: 11220,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := run(c.actual)
			assert.Equal(t, c.expected, got)
		})
	}
}

const smallops = `noop
noop
addx 2
noop
addx 3
noop`

const mediumops = `addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1`
