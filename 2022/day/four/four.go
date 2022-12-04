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
	min1 := s1[0]
	max1 := s1[1]

	min2 := s2[0]
	max2 := s2[1]

	return (min1 <= min2 && max1 >= max2 ||
		min1 >= min2 && max1 <= max2)
}

/*

 */

func overlaps(s1, s2 []int) bool {
	min1 := s1[0]
	max1 := s1[1]

	min2 := s2[0]
	max2 := s2[1]

	return (min1 >= min2 && min1 <= max2) ||
		(max1 >= min2 && max1 <= max2) ||
		(min2 >= min1 && min2 <= max1) ||
		(max2 >= min1 && max2 <= max1)
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
