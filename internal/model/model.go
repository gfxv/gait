package model

type Model interface {
	GetSuggestion(string) (string, error)
}
