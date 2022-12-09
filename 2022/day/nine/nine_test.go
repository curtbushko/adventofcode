package nine

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
			if c.expected != got {
				t.Errorf("expected %v, got %v", c.expected, got)
			}
		})
	}
}

func Test_RunTwo(t *testing.T) {
	cases := []struct {
		name     string
		actual   string
		expected int
	}{
		{
			name:     "example",
			actual:   "larger_example.input",
			expected: 36,
		},
		// {
		// 	name:     "large example",
		// 	actual:   "large.input",
		// 	expected: 6175,
		// },
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := runTwo(c.actual)
			if c.expected != got {
				t.Errorf("expected %v, got %v", c.expected, got)
			}
		})
	}
}

func Test_Moves(t *testing.T) {

	head := Head{
		Coords{x: 0, y: 0},
	}
	tail := Tail{
		Coords{x: 0, y: 0},
	}

	// == R 4 ==
	head.Move("R")
	assert.Equal(t, 1, head.x)
	assert.Equal(t, 0, head.y)

	tail.Move(head)
	assert.Equal(t, 0, tail.x)
	assert.Equal(t, 0, tail.y)

	head.Move("R")
	assert.Equal(t, 2, head.x)
	assert.Equal(t, 0, head.y)

	tail.Move(head)
	assert.Equal(t, 1, tail.x)
	assert.Equal(t, 0, tail.y)

	head.Move("R")
	assert.Equal(t, 3, head.x)
	assert.Equal(t, 0, head.y)

	tail.Move(head)
	assert.Equal(t, 2, tail.x)
	assert.Equal(t, 0, tail.y)

	head.Move("R")
	assert.Equal(t, 4, head.x)
	assert.Equal(t, 0, head.y)

	tail.Move(head)
	assert.Equal(t, 3, tail.x)
	assert.Equal(t, 0, tail.y)

	// == U 4 ==
	head.Move("U")
	assert.Equal(t, 4, head.x)
	assert.Equal(t, 1, head.y)

	tail.Move(head)
	assert.Equal(t, 3, tail.x)
	assert.Equal(t, 0, tail.y)

	head.Move("U")
	assert.Equal(t, 4, head.x)
	assert.Equal(t, 2, head.y)

	tail.Move(head)
	assert.Equal(t, 4, tail.x)
	assert.Equal(t, 1, tail.y)

}
