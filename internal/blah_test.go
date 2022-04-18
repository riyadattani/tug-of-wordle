package internal_test

import (
	"reflect"
	"testing"

	"tug-of-wordle/internal"
)

func TestName(t *testing.T) {
	t.Run("get game count, name and score", func(t *testing.T) {
		wordleOutput := `Wordle 301 3/6

⬜⬜⬜🟨🟨
🟩🟩🟩⬜⬜
🟩🟩🟩🟩🟩`

		name := "Riya"

		row, err := internal.GetRow(name, wordleOutput)
		if err != nil {
			t.Errorf("error getting score: %v", err)
		}

		expectedRow := internal.Row{
			Game:  "301",
			Score: internal.Score{name: "3"},
		}

		if reflect.DeepEqual(row, expectedRow) == false {
			t.Errorf("Expected row to be %v, got %v", expectedRow, row)
		}
	})

}
