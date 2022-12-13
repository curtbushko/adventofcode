package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	rounds int = 10000
)

type Monkey struct {
	id       int
	items    []int
	op       string
	test     int
	ttarget  int
	ftarget  int
	activity int
}

func (m *Monkey) IsEmpty() bool {
	return len(m.items) == 0
}

func (m Monkey) Len() int {
	return len(m.items)
}

func (m *Monkey) Throw() int {
	if m.IsEmpty() {
		return 0
	}
	i := m.items[0]
	m.items = m.items[1:]
	return i
}

func (m *Monkey) Catch(v int) {
	m.items = append(m.items, v)
}

func (m Monkey) String() string {
	return fmt.Sprintf("Monkey %d\n items: %d\n operations: %s\n test: %d\n \ttrue target: %d\n\tfalse target: %d\n", m.id, m.items, m.op, m.test, m.ttarget, m.ftarget)
}

func (m Monkey) Operation(old, reduction int) (worry int) {
	f := strings.Fields(m.op)

	old = old % reduction

	term := 0
	switch f[2] {
	case "old":
		term = old
	default:
		term, _ = strconv.Atoi(f[2])
	}

	switch f[1] {
	case "*":
		worry = old * term
	case "+":
		worry = old + term
	}

	return worry
	// return worry
}

func Reducer(m []Monkey) int {
	items := make([]int, 0)
	for _, monkies := range m {
		items = append(items, monkies.items...)
	}
	fmt.Println("items:", items)
	lcm := LCM(items[0], items[1], items...)
	return lcm
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func (m Monkey) Target(worry int) int {
	if (worry % m.test) == 0 {
		return m.ttarget
	}
	return m.ftarget
}

func main() {
	fmt.Println("hello world")
}

func makeMonkies(input string) []Monkey {
	m := make([]Monkey, 0)

	for _, v := range strings.Split(input, "\n\n") {
		m = append(m, makeMonkey(v))
	}
	return m
}

func makeMonkey(input string) Monkey {
	m := Monkey{
		items: make([]int, 0),
	}

	for i, line := range strings.Split(input, "\n") {
		f := strings.Fields(strings.ReplaceAll(line, ":", ""))
		switch i {
		case 0:
			m.id, _ = strconv.Atoi(f[1])
		case 1:
			for j, ff := range f {
				if j > 1 {
					item, _ := strconv.Atoi(strings.ReplaceAll(string(ff), ",", ""))
					m.items = append(m.items, item)
				}
			}
		case 2:
			m.op = fmt.Sprintf("%s %s %s", f[3], f[4], f[5])
		case 3:
			m.test, _ = strconv.Atoi(f[3])
		case 4:
			m.ttarget, _ = strconv.Atoi(f[5])
		case 5:
			m.ftarget, _ = strconv.Atoi(f[5])
		}
	}
	return m
}

func processMonkeys(m []Monkey) (total int) {
	activity := make([]int, len(m))

	for i := 0; i < rounds; i++ {
		// At the beginning of each round reduce the worry # for each monkey using LCM
		reduction := Reducer(m)
		fmt.Println("Reduction: ", reduction)

		for thrower := range m {
			num := len(m[thrower].items)
			for j := 0; j < num; j++ {
				worry := m[thrower].Operation(m[thrower].items[0], reduction)
				target := m[thrower].Target(worry)
				// fmt.Println("Monkey ", m[thrower].id, "items: ", m[thrower].items, "worry: ", worry, "target: ", target)
				_ = m[thrower].Throw()
				activity[thrower]++
				m[target].Catch(worry)
				// activity[target]++
			}
		}

		// for thrower := range m {
		// 	fmt.Println("Round: ", i, "Monkey: ", m[thrower].id, "items: ", m[thrower].items)
		// }
	}

	// business := make([]int, 0)
	for k := range activity {
		fmt.Println("Activity for Monkey: ", k, "is: ", activity[k])
	}
	sort.Ints(activity)
	fmt.Println(activity)
	len := len(activity)
	total = activity[len-1] * activity[len-2]
	fmt.Println("Total:", total)
	return total
}

func run(filename string) int {
	input, _ := os.Open(filename)
	defer input.Close()
	sc := bufio.NewScanner(input)

	for sc.Scan() {
		fmt.Println(sc.Text())
	}
	return 0
}
