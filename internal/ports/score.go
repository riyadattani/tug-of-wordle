package ports

type Wordle interface {
	GetNameAndScore() (string, string, error)
}
