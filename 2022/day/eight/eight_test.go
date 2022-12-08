package eight

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ProcessGrid(t *testing.T) {
	cases := []struct {
		name     string
		actual   string
		expected int
	}{
		{
			name:     "example",
			actual:   "example.input",
			expected: 21,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := ProcessGrid(c.actual)
			assert.NoError(t, err)
			assert.Equal(t, c.expected, got)
		})
	}
}

func Test_CurrentLine(t *testing.T) {
	g := &Grid{
		previous: "30373",
		current:  "25512",
		next:     "65332",
	}

	got := g.ProcessCurrentLine()
	assert.Equal(t, 4, got)
}
