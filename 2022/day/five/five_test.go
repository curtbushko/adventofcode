package five

import (
	"testing"
)

const (
	example_header = ` 1   2   3 `
	large_header   = ` 1   2   3   4   5   6   7   8   9 `
)

func Test_Run(t *testing.T) {
	cases := []struct {
		name      string
		filename  string
		header    string
		stackSize int
		part      int
		expected  string
	}{
		{
			name:      "Part 1: Example input",
			filename:  "example.input",
			header:    example_header,
			stackSize: 3,
			part:      1,
			expected:  "CMZ",
		},
		{
			name:      "Part 1: Large input",
			filename:  "large.input",
			header:    large_header,
			stackSize: 9,
			part:      1,
			expected:  "PTWLTDSJV",
		},
		{
			name:      "Part 2: Large input",
			filename:  "large.input",
			header:    large_header,
			stackSize: 9,
			part:      2,
			expected:  "WZMFVGGZP",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := run(c.filename, c.header, c.stackSize, c.part)
			if c.expected != got {
				t.Errorf("expected %v, got %v", c.expected, got)
			}
		})
	}
}
