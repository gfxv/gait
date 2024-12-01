package model

import "os"

var prompt string

type Model interface {
	GetSuggestion(string) (string, error)
}

func LoadPrompt(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	prompt = string(data)
	return nil
}
