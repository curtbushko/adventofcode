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
	rounds int = 2
)

type Troop struct {
	m []Monkey
}

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

func (m Monkey) Operation(old int) (worry int) {
	f := strings.Fields(m.op)

	term := 0
	switch f[1] {
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

	return worry / 3
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

func makeTroop(input string) Troop {
	t := Troop{
		m: make([]Monkey, 0),
	}

	for _, v := range strings.Split(input, "\n\n") {
		t.m = append(t.m, makeMonkey(v))
	}
	return t
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

func (t *Troop) processMonkeys() (total int) {
	for i := 0; i < rounds; i++ {
		fmt.Println("Round: ", i)
		for _, mon := range t.m {
			num := len(mon.items)
			for j := 0; j < num; j++ {
				worry := mon.Operation(mon.items[0])
				target := mon.Target(worry)
				it := mon.Throw()
				// fmt.Println("Monkey ", m.id, "throws ", it, "to ", "Monkey ", target)
				t.m[target].Catch(it)
				fmt.Println("Target: ", target, "Target items:", t.m[target].items)
				mon.activity++
				t.m[target].activity++
				// fmt.Println("Round: ", i, "Monkey: ", mon.id, "items: ", mon.items)
			}
		}

	}

	business := make([]int, 0)
	for k, m := range t.m {
		fmt.Println("Activity for Monkey: ", k, "is: ", m.activity)
		business = append(business, m.activity)
	}
	sort.Ints(business)
	len := len(business)
	total = business[len-1] * business[len-2]
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
