package three

import "testing"

func Test_newRucksack(t *testing.T) {
	cases := []struct {
		name         string
		actual       string
		expected_one string
		expected_two string
	}{
		{
			name:         "example",
			actual:       "vJrwpWtwJgWrhcsFMMfFFhFp",
			expected_one: "vJrwpWtwJgWr",
			expected_two: "hcsFMMfFFhFp",
		},
		{
			name:         "example two",
			actual:       "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			expected_one: "jqHRNqRjqzjGDLGL",
			expected_two: "rsFMfFZSrLrFZsSL",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r := newRucksack(c.actual)
			if c.expected_one != r.first {
				t.Errorf("expected one %v, got one %v", c.expected_one, r.first)
			}

			if c.expected_two != r.second {
				t.Errorf("expected two %v, got two %v", c.expected_two, r.second)
			}
		})
	}
}

func Test_FindDuplicate(t *testing.T) {
	cases := []struct {
		name     string
		actual   string
		expected rune
	}{
		{
			name:     "example",
			actual:   "vJrwpWtwJgWrhcsFMMfFFhFp",
			expected: rune('p'),
		},
		{
			name:     "example two",
			actual:   "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			expected: rune('L'),
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r := newRucksack(c.actual)
			r.findDuplicate()
			if c.expected != r.duplicate {
				t.Errorf("expected %v, got %v", c.expected, r.duplicate)
			}
		})
	}
}

func Test_Value(t *testing.T) {
	cases := []struct {
		name     string
		actual   rune
		expected int
	}{
		{
			name:     "a",
			actual:   rune('a'),
			expected: 1,
		},
		{
			name:     "z",
			actual:   rune('z'),
			expected: 26,
		},
		{
			name:     "A",
			actual:   rune('A'),
			expected: 27,
		},
		{
			name:     "Z",
			actual:   rune('Z'),
			expected: 52,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := value(c.actual)
			if c.expected != got {
				t.Errorf("expected %v, got %v", c.expected, got)
			}
		})
	}
}

func Test_FindBadge(t *testing.T) {
	cases := []struct {
		name     string
		s1       string
		s2       string
		s3       string
		expected rune
	}{
		{
			name:     "example",
			s1:       "vJrwpWtwJgWrhcsFMMfFFhFp",
			s2:       "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			s3:       "PmmdzqPrVvPwwTWBwg",
			expected: rune('r'),
		},
		{
			name: "example 2",

			s1:       "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			s2:       "ttgJtRGJQctTZtZT",
			s3:       "CrZsJsPPZsGzwwsLwLmpwMDw",
			expected: rune('Z'),
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := findBadge(c.s1, c.s2, c.s3)
			if c.expected != got {
				t.Errorf("expected %v, got %v", c.expected, got)
			}
		})
	}
}

func Test_Packing(t *testing.T) {
	cases := []struct {
		name     string
		actual   string
		expected int
	}{
		{
			name:     "example",
			actual:   input,
			expected: 70,
		},
		{
			name:     "example",
			actual:   input2,
			expected: 2641,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := packing(c.actual)
			if c.expected != got {
				t.Errorf("expected %v, got %v", c.expected, got)
			}
		})
	}
}
