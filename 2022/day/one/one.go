package one

import (
	"regexp"
	"strconv"
	"strings"
)

func MaxCalories(m map[int]string) (index, max int) {
	var total int
	var j int

	for i, v := range m {
		total = 0
		fields := strings.Fields(v)
		for _, element := range fields {
			j, _ = strconv.Atoi(element)
			total = total + j
		}
		if total > max {
			max = total
			index = i
		}
	}
	return index, max
}

func TopThree(in string) int {
	var index int
	var max int
	var total int

	m := Splitter(in)

	index, max = MaxCalories(m)
	total = max
	delete(m, index)

	index, max = MaxCalories(m)
	total = total + max
	delete(m, index)

	index, max = MaxCalories(m)
	total = total + max
	delete(m, index)

	return total
}

func Splitter(str string) map[int]string {
	m := make(map[int]string)
	re := regexp.MustCompile("\n")
	for i, line := range strings.Split(str, "\n\n") {
		line = re.ReplaceAllString(line, " ")
		m[i] = line
	}
	return m
}
