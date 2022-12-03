package two

import (
	"strings"
)

const (
	lose int = 0
	draw int = 3
	win  int = 6

	rock     int = 1
	paper    int = 2
	scissors int = 3
)

type round struct {
	opponent int
	player   int
}

/*
****	player(x)  rock   paper   scissors
opponent(y))
rock     	      draw   win     lose
paper           win    draw    lose
scissors        lose   win     draw
*/
func (r round) result() int {
	matrix := [3][3]int{
		{draw, win, lose},
		{lose, draw, win},
		{win, lose, draw},
	}
	return matrix[r.opponent-1][r.player-1] + r.player
}

func strategy(opponent, player int) int {
	// lose int = 0
	// draw int = 3
	// win  int = 6

	m := map[int]map[int]int{
		lose: {
			rock:     scissors,
			paper:    rock,
			scissors: paper,
		},
		draw: {
			rock:     rock,
			paper:    paper,
			scissors: scissors,
		},
		win: {
			rock:     paper,
			paper:    scissors,
			scissors: rock,
		},
	}

	return m[player][opponent]
}

/*
A = rock
B = paper
C = scissors

X = lose
Y = draw
Z = win
*/
func NewRound(line string) round {
	m := map[string]int{
		"X": lose,
		"Y": draw,
		"Z": win,
		"A": rock,
		"B": paper,
		"C": scissors,
	}

	c := make(map[int]string)
	for i, move := range strings.Split(line, " ") {
		c[i] = move
	}

	player := strategy(m[c[0]], m[c[1]])

	return round{
		opponent: m[c[0]],
		player:   player,
	}
}

func tournament(input string) int {
	total := 0
	var r round
	for _, line := range strings.Split(input, "\n") {
		r = NewRound(line)
		total = total + r.result()
	}

	return total
}
