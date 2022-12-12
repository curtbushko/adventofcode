package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MakeMonkey(t *testing.T) {
	cases := []struct {
		name     string
		actual   string
		expected Monkey
	}{
		{
			name:   "example",
			actual: monkey0,
			expected: Monkey{
				id:       0,
				items:    []int{79, 98},
				op:       "old * 19",
				test:     23,
				ttarget:  0,
				ftarget:  0,
				activity: 0,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := makeMonkey(c.actual)
			assert.Equal(t, c.expected, got)
		})
	}
}

func Test_MakeMonkies(t *testing.T) {
	cases := []struct {
		name     string
		actual   string
		expected []Monkey
	}{
		{
			name:   "example",
			actual: monkey0,
			expected: []Monkey{
				{
					id:       0,
					items:    []int{79, 98},
					op:       "old * 19",
					test:     23,
					ttarget:  0,
					ftarget:  0,
					activity: 0,
				},
			},
		},
		{
			name:   "example input from aoc",
			actual: inputaoc,
			expected: []Monkey{
				{
					id:       0,
					items:    []int{79, 98},
					op:       "old * 19",
					test:     23,
					ttarget:  2,
					ftarget:  3,
					activity: 0,
				},
				{
					id:       1,
					items:    []int{54, 65, 75, 74},
					op:       "old + 6",
					test:     19,
					ttarget:  2,
					ftarget:  0,
					activity: 0,
				},
				{
					id:       2,
					items:    []int{79, 60, 97},
					op:       "old * old",
					test:     13,
					ttarget:  1,
					ftarget:  3,
					activity: 0,
				},
				{
					id:       3,
					items:    []int{74},
					op:       "old + 3",
					test:     17,
					ttarget:  0,
					ftarget:  1,
					activity: 0,
				},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := makeMonkies(c.actual)
			assert.Equal(t, c.expected, got)
		})
	}
}

func Test_ProcessMonkeys(t *testing.T) {
	cases := []struct {
		name     string
		actual   string
		expected int
	}{
		{
			name:     "example",
			actual:   inputexample,
			expected: 10605,
		},

		// {
		// 	name:   "large input",
		// 	actual: largeinput,
		//
		// 	// expected: 15310845153,
		// 	expected: 72884,
		// },

		// large expected (from file) 15310845153

	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			m := makeMonkies(c.actual)
			got := processMonkeys(m)
			assert.Equal(t, c.expected, got)
		})
	}
}

func Test_MonkeyOperation(t *testing.T) {
	cases := []struct {
		name     string
		actual   string
		expected int
	}{
		// {
		// 	name:     "monkey 0",
		// 	actual:   monkey0,
		// 	expected: 500,
		// },
		// {
		// 	name:     "monkey 1",
		// 	actual:   monkey1,
		// 	expected: 20,
		// },
		{
			name:     "monkey 2",
			actual:   monkey2,
			expected: 6241,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			monkey := makeMonkey(c.actual)
			got := monkey.Operation(monkey.items[0])
			assert.Equal(t, c.expected, got)
		})
	}
}

func Test_MonkeyTarget(t *testing.T) {
	cases := []struct {
		name     string
		actual   string
		expected int
	}{
		{
			name:     "monkey 0",
			actual:   monkey0,
			expected: 0,
		},
		{
			name:     "monkey 1",
			actual:   monkey1,
			expected: 0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			monkey := makeMonkey(c.actual)
			worry := monkey.Operation(monkey.items[0])
			got := monkey.Target(worry)
			assert.Equal(t, c.expected, got)
		})
	}
}

func Test_ThrowCatch(t *testing.T) {
	m := []Monkey{
		{
			id:    0,
			items: []int{79, 98},
		},
		{
			id:    1,
			items: []int{54, 65, 75, 74},
		},
	}

	assert.Equal(t, 2, m[0].Len(), "expected monkey length to be 0")

	m[0].Catch(1)
	m[0].Catch(2)
	m[0].Catch(3)

	m[1].Catch(4)
	m[1].Catch(5)
	m[1].Catch(6)

	// Verify that the monkey is not empty
	assert.False(t, m[0].IsEmpty(), "expected monkey to be not empty")
	assert.False(t, m[1].IsEmpty(), "expected monkey to be not empty")

	// Verify that the length of the monkey is 3
	assert.Equal(t, 5, m[0].Len(), "expected monkey length to be 5")

	i := m[0].Throw()
	assert.Equal(t, 79, i, "expected item to be 79")

	assert.Equal(t, 7, m[1].Len(), "expected monkey length to be 7")
	m[1].Catch(i)
	assert.Equal(t, 8, m[1].Len(), "expected monkey length to be 8")
}

const monkey0 = `Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 0 
    If false: throw to monkey 0`

const monkey1 = `Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0`

const monkey2 = `Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3`

const inputaoc = `Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1`

const inputexample = `Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1`

const largeinput = `Monkey 0:
  Starting items: 89, 95, 92, 64, 87, 68
  Operation: new = old * 11
  Test: divisible by 2
    If true: throw to monkey 7
    If false: throw to monkey 4

Monkey 1:
  Starting items: 87, 67
  Operation: new = old + 1
  Test: divisible by 13
    If true: throw to monkey 3
    If false: throw to monkey 6

Monkey 2:
  Starting items: 95, 79, 92, 82, 60
  Operation: new = old + 6
  Test: divisible by 3
    If true: throw to monkey 1
    If false: throw to monkey 6

Monkey 3:
  Starting items: 67, 97, 56
  Operation: new = old * old
  Test: divisible by 17
    If true: throw to monkey 7
    If false: throw to monkey 0

Monkey 4:
  Starting items: 80, 68, 87, 94, 61, 59, 50, 68
  Operation: new = old * 7
  Test: divisible by 19
    If true: throw to monkey 5
    If false: throw to monkey 2

Monkey 5:
  Starting items: 73, 51, 76, 59
  Operation: new = old + 8
  Test: divisible by 7
    If true: throw to monkey 2
    If false: throw to monkey 1

Monkey 6:
  Starting items: 92
  Operation: new = old + 5
  Test: divisible by 11
    If true: throw to monkey 3
    If false: throw to monkey 0

Monkey 7:
  Starting items: 99, 76, 78, 76, 79, 90, 89
  Operation: new = old + 7
  Test: divisible by 5
    If true: throw to monkey 4
    If false: throw to monkey 5`
