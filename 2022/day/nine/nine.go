package nine

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const knots = 9

type Coords struct {
	x int
	y int
}

type Head struct {
	Coords
}

type Tail struct {
	Coords
}

func run(filename string) int {
	input, _ := os.Open(filename)
	defer input.Close()
	sc := bufio.NewScanner(input)

	m := make(map[Coords]bool)
	h := Head{
		Coords{x: 0, y: 0},
	}
	t := Tail{
		Coords{x: 0, y: 0},
	}
	var dir string
	var dist int

	for sc.Scan() {
		parts := strings.Split(sc.Text(), " ")
		dir = parts[0]
		dist, _ = strconv.Atoi(parts[1])
		for i := 0; i < dist; i++ {
			h.Move(dir)
			t.Move(h)
			m[Coords{t.x, t.y}] = true
		}

	}
	// Count the places touched by Y
	total := 0
	for _, v := range m {
		if v == true {
			total++
		}
	}
	fmt.Println(total)
	return total
}

func runTwo(filename string) int {
	input, _ := os.Open(filename)
	defer input.Close()
	sc := bufio.NewScanner(input)

	m := make(map[Coords]bool)
	h := Head{
		Coords{x: 0, y: 0},
	}
	t := make([]Tail, knots)

	var dir string
	var dist int

	for sc.Scan() {
		parts := strings.Split(sc.Text(), " ")
		dir = parts[0]
		dist, _ = strconv.Atoi(parts[1])
		for i := 0; i < dist; i++ {
			h.Move(dir)
			t[0].Move(h)
			m[Coords{t[0].x, t[0].y}] = true
			for j := 1; j < knots; j++ {
				t[j].MoveTail(t[j-1])
				m[Coords{t[0].x, t[0].y}] = true
			}
		}

	}
	// Count the places touched by Y
	total := 0
	for _, v := range m {
		if v == true {
			total++
		}
	}
	fmt.Println(total)
	return total
}

func (h *Head) Move(dir string) {
	switch dir {
	case "U":
		h.y++
	case "D":
		h.y--
	case "L":
		h.x--
	case "R":
		h.x++
	}
}

func (t *Tail) Move(h Head) {
	// Get the absolute distance to see if we need to move or not
	distX := math.Abs(float64(t.x - h.x))
	distY := math.Abs(float64(t.y - h.y))
	if distX <= float64(1) && distY <= float64(1) {
		return
	}

	dirX := 1
	dirY := 1

	if h.y < t.y {
		dirY = -1
	}
	if h.x < t.x {
		dirX = -1
	}

	if h.x == t.x { // they are on the same row
		t.y += 1 * dirY
		return
	}

	if h.y == t.y { // they are the same column
		t.x += 1 * dirX
		return
	}

	// stupid diagonal
	t.y += 1 * dirY
	t.x += 1 * dirX
}

func (t *Tail) MoveTail(t2 Tail) {
	// Get the absolute distance to see if we need to move or not
	distX := math.Abs(float64(t.x - t2.x))
	distY := math.Abs(float64(t.y - t2.y))
	if distX <= float64(1) && distY <= float64(1) {
		return
	}

	dirX := 1
	dirY := 1

	if t2.y < t.y {
		dirY = -1
	}
	if t2.x < t.x {
		dirX = -1
	}

	if t2.x == t.x { // they are on the same row
		t.y += 1 * dirY
		return
	}

	if t2.y == t.y { // they are the same column
		t.x += 1 * dirX
		return
	}

	// stupid diagonal
	t.y += 1 * dirY
	t.x += 1 * dirX
}
