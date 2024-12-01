package main

import (
	"flag"
	"gait/internal/app"
	"gait/internal/model"

	"github.com/joho/godotenv"

	"log"
	"os"
)

const envPath = "local.env"
const modelName = "gemini-1.5-flash"

var promptPath string

func init() {
	flag.StringVar(&promptPath, "prompt", "", "Path to the prompt file")
	flag.Parse()

	model.LoadPrompt(promptPath)
}

func main() {

	if err := godotenv.Load(envPath); err != nil {
		log.Fatal(err)
	}

	apiKey := os.Getenv("GAIT_API_KEY")
	if len(apiKey) == 0 {
		log.Fatal("No API key found")
	}

	gemini := model.NewGeminiModel(apiKey, modelName)
	cli := app.NewApp(gemini)
	cli.MustRun()
}
