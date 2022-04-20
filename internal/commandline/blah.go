package commandline

import (
	"flag"
	"os"
)

func GetValues() (string, string) {
	namePtr := flag.String("name", "", "e.g. Susan")
	scorePtr := flag.String("score", "", "paste your wordle score")
	flag.Parse()

	if *namePtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *scorePtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	return *namePtr, *scorePtr
}
