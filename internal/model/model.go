package model

import (
	"context"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

const PROMPT = "Create well-formed git commit message based of off the currently staged file\ncontents. The message should convey why something was changed and not what\nchanged. Use the well known format that has the prefix chore, fix, etc.\n\nOnly include changes to source files for the programming languages, shell configurations files,\ndocumentation such as readme and other .mds, and any changes to package management file.\nExclude any lock or sum files.\n\nDo not use markdown format for the output.\n\nFor the first line of the commit message, this must be constrained to 40 characters as a\nmaximum and use additional lines for any further context.\n"

type Model interface {
	GetSuggestion(string) (string, error)
}

type GeminiModel struct {
	apiKey string
	model  string
}

func NewGeminiModel(apiKey, model string) *GeminiModel {
	return &GeminiModel{
		apiKey: apiKey,
		model:  model,
	}
}

func (g *GeminiModel) GetSuggestion(diff string) (string, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(g.apiKey))
	if err != nil {
		return "", err
	}
	defer client.Close()

	model := client.GenerativeModel(g.model)
	model.SystemInstruction = genai.NewUserContent(genai.Text(PROMPT))

	resp, err := model.GenerateContent(ctx, genai.Text(diff))
	if err != nil {
		return "", err
	}

	result := ""
	for _, c := range resp.Candidates {
		if c.Content != nil {
			for _, part := range c.Content.Parts {
				result += fmt.Sprintf("%s", part)
			}
		}
	}
	return result, nil
}
