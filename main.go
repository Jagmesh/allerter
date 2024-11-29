package main

import (
	"os"

	"tg-alerter/bot"
	"tg-alerter/logger"

	"github.com/joho/godotenv"
)

var TG_TOKEN_VAR_NAME = "TG_API_BOT_TOKEN"

func main() {
	log := logger.InitLogger()

	botToken := getBotToken()
	if botToken == "" {
		log.Fatalf("No %s provided", TG_TOKEN_VAR_NAME)
	}

	myBot, err := bot.NewBot(botToken)
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
	}

	if err := myBot.Start(); err != nil {
		log.Fatalf("Error starting bot: %v", err)
	}
}

func getBotToken() string {
	log := logger.GetLogger()

	botToken := os.Getenv(TG_TOKEN_VAR_NAME)
	if botToken == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Error("Failed loading .env: ", err)
		}

		botToken = os.Getenv(TG_TOKEN_VAR_NAME)
	}

	return botToken
}
