package main

import (
	"fmt"
	"log"
	"os"

	"tug-of-wordle/internal/commandline"
	"tug-of-wordle/internal/scoreboard"
)

func main() {
	wordle := commandline.New()

	name, score, err := wordle.GetNameAndScore()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	row, _ := scoreboard.GetRow(name, score)

	fmt.Printf("Row: %v", row)
}
