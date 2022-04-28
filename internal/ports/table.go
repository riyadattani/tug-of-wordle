package ports

type Score map[string]string

type Row struct {
	Game  string
	Score Score
}

type LeagueTable interface {
	GetRow(name, wordleOutput string) (Row, error)
}
