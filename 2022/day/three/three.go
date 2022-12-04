package three

import (
	"fmt"
	"strings"
)

const (
	lowerOffset int = 96
	upperOffset int = 38
)

type rucksack struct {
	all       string
	first     string
	second    string
	three     string
	duplicate rune
	length    int
}

func newRucksack(str string) rucksack {
	r := rucksack{
		all: str,
	}
	r.length = len(r.all)

	r.first = r.all[:r.length/2]
	r.second = r.all[r.length/2:]

	return r
}

func (r *rucksack) findDuplicate() {
	f := []rune(r.first)
	s := []rune(r.second)

	// flatten the first string into a map for O(1) constant time lookup
	m1 := make(map[rune]struct{})
	for _, v := range f {
		m1[v] = struct{}{}
	}

	// iterate string2, using the O(1) lookup.
	for _, v := range s {
		if _, exists := m1[v]; exists {
			r.duplicate = v
		}
	}
}

func value(r rune) int {
	i := int(r)
	if i > lowerOffset { // lowercase
		return i - lowerOffset
	}
	return i - upperOffset
}

func findBadge(s1, s2, s3 string) rune {
	f := []rune(s1)
	s := []rune(s2)
	t := []rune(s3)

	m1 := make(map[rune]struct{})
	for _, v := range f {
		m1[v] = struct{}{}
	}

	m2 := make(map[rune]struct{})
	for _, v := range s {
		m2[v] = struct{}{}
	}

	for _, v := range t {
		if _, exists := m1[v]; exists {
			if _, exists := m2[v]; exists {
				return v
			}
		}
	}
	return 0
}

func packing(input string) int {
	total := 0
	j := 1
	var r1, r2, r3 rucksack
	for _, line := range strings.Split(input, "\n") {
		if j == 1 {
			r1 = newRucksack(line)
			j++
			continue
		}

		if j == 3 {
			j = 1
			r3 = newRucksack(line)
			fmt.Println(r1.all)
			fmt.Println(r2.all)
			fmt.Println(r3.all)
			total = total + value(findBadge(r1.all, r2.all, r3.all))
			continue
		}

		if j == 2 {
			j++
			r2 = newRucksack(line)
			continue
		}
	}

	return total
}
