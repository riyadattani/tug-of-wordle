package internal_test

import (
	"testing"

	"tug-of-wordle/internal"
)

func TestName(t *testing.T) {
	t.Run("retrieve wordle score and game number", func(t *testing.T) {
		wordleOutput := `Wordle 301 3/6

⬜⬜⬜🟨🟨
🟩🟩🟩⬜⬜
🟩🟩🟩🟩🟩`

		gotScore, err := internal.GetScore(wordleOutput)
		if err != nil {
			t.Errorf("error getting score: %v", err)
		}

		expectedScore := "3"
		if gotScore != expectedScore {
			t.Errorf("Expected score to be %v, got %v", expectedScore, gotScore)
		}
	})
}
