package main

import (
	"fmt"

	"tug-of-wordle/internal/commandline"
	"tug-of-wordle/internal/scoreboard"
)

func main() {
	name, score := commandline.GetValues()
	row, _ := scoreboard.GetRow(name, score)

	fmt.Printf("Row: %v", row)
}
