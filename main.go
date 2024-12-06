package main

import (
	"gait/internal/app"
	"gait/internal/model"

	_ "embed"
	"log"
	"os"
)

const envPath = "local.env"
const modelName = "gemini-1.5-flash"

//go:embed prompt
var prompt string

func init() {

	model.SetPrompt(prompt)
}

func main() {

	apiKey := os.Getenv("GAIT_API_KEY")
	if len(apiKey) == 0 {
		log.Fatal("No API key found")
	}

	gemini := model.NewGeminiModel(apiKey, modelName)
	cli := app.NewApp(gemini)
	cli.MustRun()
}
