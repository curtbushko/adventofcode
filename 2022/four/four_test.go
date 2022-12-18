package four

import "testing"

func Test_ExpandID(t *testing.T) {
	cases := []struct {
		name     string
		actual   string
		expected []int
	}{
		{
			name:     "2-4",
			actual:   "2-4",
			expected: []int{2, 4},
		},
		{
			name:     "8-12",
			actual:   "8-12",
			expected: []int{8, 12},
		},
		{
			name:     "18-22",
			actual:   "18-22",
			expected: []int{18, 22},
		},
		{
			name:     "52-52",
			actual:   "52-52",
			expected: []int{52, 52},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := expandID(c.actual)
			if c.expected[0] != got[0] || c.expected[1] != got[1] {
				t.Errorf("expected %v, got %v", c.expected, got)
			}
		})
	}
}

func Test_Contains(t *testing.T) {
	cases := []struct {
		name     string
		s1       []int
		s2       []int
		expected bool
	}{
		{
			name:     "1-3 does not contain 4-5",
			s1:       []int{1, 3},
			s2:       []int{4, 5},
			expected: false,
		},
		{
			name:     "2-8 contains 3-7",
			s1:       []int{2, 8},
			s2:       []int{3, 7},
			expected: true,
		},
		{
			name:     "2-8 contains 3-7 (flipped)",
			s1:       []int{3, 7},
			s2:       []int{2, 8},
			expected: true,
		},

		{
			name:     "4-6 contains 6-6",
			s1:       []int{4, 6},
			s2:       []int{6, 6},
			expected: true,
		},
		{
			name:     "4-6 contains 6-6 (flipped)",
			s1:       []int{6, 6},
			s2:       []int{4, 6},
			expected: true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := contains(c.s1, c.s2)
			if c.expected != got {
				t.Errorf("expected %v, got %v", c.expected, got)
			}
		})
	}
}

func Test_Overlaps(t *testing.T) {
	cases := []struct {
		name     string
		s1       []int
		s2       []int
		expected bool
	}{
		{
			name:     "1-3 does not overlap 4-5",
			s1:       []int{1, 3},
			s2:       []int{4, 5},
			expected: false,
		},
		{
			name:     "2-8 overlaps 3-7",
			s1:       []int{2, 8},
			s2:       []int{3, 7},
			expected: true,
		},
		{
			name:     "2-8 overlaps 3-7 (flipped)",
			s1:       []int{3, 7},
			s2:       []int{2, 8},
			expected: true,
		},

		{
			name:     "4-6 overlaps 6-6",
			s1:       []int{4, 6},
			s2:       []int{6, 6},
			expected: true,
		},
		{
			name:     "4-6 overlaps 6-6 (flipped)",
			s1:       []int{6, 6},
			s2:       []int{4, 6},
			expected: true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := overlaps(c.s1, c.s2)
			if c.expected != got {
				t.Errorf("expected %v, got %v", c.expected, got)
			}
		})
	}
}

func Test_CleanupContains(t *testing.T) {
	cases := []struct {
		name     string
		actual   string
		expected int
	}{
		{
			name:     "example",
			actual:   input,
			expected: 2,
		},
		{
			name:     "example 2",
			actual:   input2,
			expected: 536,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := cleanupContains(c.actual)
			if c.expected != got {
				t.Errorf("expected %v, got %v", c.expected, got)
			}
		})
	}
}

func Test_CleanupOverlaps(t *testing.T) {
	cases := []struct {
		name     string
		actual   string
		expected int
	}{
		{
			name:     "example",
			actual:   input,
			expected: 4,
		},
		{
			name:     "example 2",
			actual:   input2,
			expected: 845,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := cleanupOverlaps(c.actual)
			if c.expected != got {
				t.Errorf("expected %v, got %v", c.expected, got)
			}
		})
	}
}
