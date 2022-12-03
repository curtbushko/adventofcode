package two

import "testing"

func TestMatrix(t *testing.T) {
	cases := []struct {
		name     string
		actual   round
		expected int
	}{
		{
			name: "rock vs rock = draw",
			actual: round{
				opponent: rock,
				player:   rock,
			},
			expected: draw + rock,
		},
		{
			name: "rock vs paper = win",
			actual: round{
				opponent: rock,
				player:   paper,
			},
			expected: win + paper,
		},
		{
			name: "rock vs scissors = lose",
			actual: round{
				opponent: rock,
				player:   scissors,
			},
			expected: lose + scissors,
		},
		{
			name: "paper vs scissors = win",
			actual: round{
				opponent: paper,
				player:   scissors,
			},
			expected: win + scissors,
		},
		{
			name: "scissors vs scissors = draw",
			actual: round{
				opponent: scissors,
				player:   scissors,
			},
			expected: draw + scissors,
		},
		{
			name: "scissors vs rock = win",
			actual: round{
				opponent: scissors,
				player:   rock,
			},
			expected: win + rock,
		},
		{
			name: "paper vs rock = lose",
			actual: round{
				opponent: paper,
				player:   rock,
			},
			expected: lose + rock,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.actual.result()
			if c.expected != got {
				t.Errorf("expected %v, got %v", c.expected, got)
			}
		})
	}
}

func TestNewRound(t *testing.T) {
	cases := []struct {
		name     string
		actual   string
		expected round
	}{
		{
			name:   "A Y = rock draw = rock",
			actual: "A Y",
			expected: round{
				opponent: rock,
				player:   rock,
			},
		},
		{
			name:   "B X = paper lose = rock",
			actual: "B X",
			expected: round{
				opponent: paper,
				player:   rock,
			},
		},
		{
			name:   "C Z = scissors win = rock",
			actual: "C Z",
			expected: round{
				opponent: scissors,
				player:   rock,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := newRound(c.actual)
			if c.expected != got {
				t.Errorf("expected %v, got %v", c.expected, got)
			}
		})
	}
}

func TestStrategy(t *testing.T) {
	cases := []struct {
		name     string
		opponent int
		player   int
		expected int
	}{
		{
			name:     "opponent: rock and lose",
			opponent: rock,
			player:   lose,
			expected: scissors,
		},
		{
			name:     "opponent: rock and win",
			opponent: rock,
			player:   win,
			expected: paper,
		},
		{
			name:     "opponent: scissors and win",
			opponent: scissors,
			player:   win,
			expected: rock,
		},
		{
			name:     "opponent: scissors and draw",
			opponent: scissors,
			player:   draw,
			expected: scissors,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := strategy(c.opponent, c.player)
			if c.expected != got {
				t.Errorf("expected %v, got %v", c.expected, got)
			}
		})
	}
}

func TestTournament(t *testing.T) {
	cases := []struct {
		name     string
		actual   string
		expected int
	}{
		{
			name:     "example",
			actual:   input,
			expected: 12,
		},
		{
			name:     "example 2",
			actual:   input2,
			expected: 13131,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := tournament(c.actual)
			if c.expected != got {
				t.Errorf("expected %v, got %v", c.expected, got)
			}
		})
	}
}
