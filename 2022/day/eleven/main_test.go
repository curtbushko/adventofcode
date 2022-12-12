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

func Test_MakeTroop(t *testing.T) {
	cases := []struct {
		name     string
		actual   string
		expected Troop
	}{
		{
			name:   "example",
			actual: monkey0,
			expected: Troop{
				[]Monkey{
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
		},
		{
			name:   "example input from aoc",
			actual: inputaoc,
			expected: Troop{
				[]Monkey{
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
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := makeTroop(c.actual)
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
		// 	name:     "example",
		// 	actual:   inputaoc,
		// 	expected: 72884,
		// },
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			troop := makeTroop(c.actual)
			got := troop.processMonkeys()
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
		{
			name:     "monkey 0",
			actual:   monkey0,
			expected: 500,
		},
		{
			name:     "monkey 1",
			actual:   monkey1,
			expected: 20,
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
