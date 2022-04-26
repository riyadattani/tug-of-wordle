package scoreboard

import (
	"strings"

	"tug-of-wordle/internal/ports"
)

func GetRow(name, wordleOutput string) (ports.Row, error) {
	s := strings.Split(wordleOutput, " ")
	value := s[2][:1]
	gameNumber := s[1]
	score := ports.Score{name: value}

	return ports.Row{Game: gameNumber, Score: score}, nil
}
