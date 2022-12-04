package four

import (
	"strconv"
	"strings"
)

func expandID(str string) []int {
	s := strings.Split(str, "-")

	r1, _ := strconv.Atoi(s[0])
	r2, _ := strconv.Atoi(s[1])
	return []int{r1, r2}
}

func contains(s1, s2 []int) bool {
	// ie 1-2 and 3-4
	if s1[0] < s2[0] && s1[1] < s2[1] {
		return false
	}

	// ie 3-4 and 1-2
	if s1[0] > s2[0] && s1[1] > s2[1] {
		return false
	}

	if s1[0] <= s2[0] && s1[1] >= s2[1] {
		return true
	}

	if s1[0] >= s2[0] && s1[1] <= s2[1] {
		return true
	}
	return false
}

func overlaps(s1, s2 []int) bool {
	m1 := make(map[int]struct{})
	for i := s1[0]; i <= s1[1]; i++ {
		m1[i] = struct{}{}
	}

	m2 := make(map[int]struct{})
	for i := s2[0]; i <= s2[1]; i++ {
		m2[i] = struct{}{}
	}

	for i := range m1 {
		if _, exists := m2[i]; exists {
			return true
		}
	}

	for i := range m2 {
		if _, exists := m1[i]; exists {
			return true
		}
	}
	return false
}

func cleanupContains(input string) int {
	total := 0
	var ids []string
	for _, line := range strings.Split(input, "\n") {
		ids = strings.Split(line, ",")
		if contains(expandID(ids[0]), expandID(ids[1])) {
			total++
		}
	}
	return total
}

func cleanupOverlaps(input string) int {
	total := 0
	var ids []string
	for _, line := range strings.Split(input, "\n") {
		ids = strings.Split(line, ",")
		if overlaps(expandID(ids[0]), expandID(ids[1])) {
			total++
		}
	}
	return total
}
