package commandline

import (
	"flag"
	"fmt"
)

type CommandLine struct {
}

func New() *CommandLine {
	return &CommandLine{}
}

func (c CommandLine) GetNameAndScore() (string, string, error) {
	scorePtr := flag.String("score", "", "paste your wordle score")
	namePtr := flag.String("name", "", "e.g. Susan")
	flag.Parse()

	if *namePtr == "" {
		flag.PrintDefaults()
		return "", "", fmt.Errorf("error parsing name")
	}

	if *scorePtr == "" {
		flag.PrintDefaults()
		return "", "", fmt.Errorf("error parsing score")
	}

	return *namePtr, *scorePtr, nil
}
