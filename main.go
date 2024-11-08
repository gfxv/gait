package main

import (
	"gait/internal/app"
	"gait/internal/model"
	"github.com/joho/godotenv"

	"log"
	"os"
)

const envPath = "local.env"
const modelName = "gemini-1.5-flash"

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
