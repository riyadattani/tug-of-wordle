package internal

import "strings"

func GetScore(wordleOutput string) (string, error) {
	s := strings.Split(wordleOutput, " ")
	score := s[2][:1]
	return score, nil
}
