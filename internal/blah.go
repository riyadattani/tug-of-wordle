package internal

import "strings"

type Score map[string]string

type Row struct {
	Game  string
	Score Score
}

// {'300': [{'Riya': '2'}, {'Adi': '3'}]}

func GetRow(name, wordleOutput string) (Row, error) {
	s := strings.Split(wordleOutput, " ")
	value := s[2][:1]
	gameNumber := s[1]
	score := Score{name: value}

	return Row{gameNumber, score}, nil

}
