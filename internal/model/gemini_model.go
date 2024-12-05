package model

import (
	"context"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

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
	model.SystemInstruction = genai.NewUserContent(genai.Text(prompt))

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
