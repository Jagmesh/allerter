package bot

import (
	"tg-alerter/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	API *tgbotapi.BotAPI
}

func NewBot(token string) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	log := logger.GetLogger()
	tgbotapi.SetLogger(log)
	api.Debug = true
	log.Info("Authorized as: ", api.Self.UserName)

	return &Bot{API: api}, nil
}

func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.API.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		go b.HandleUpdate(update) // Обрабатываем каждый апдейт в отдельной горутине
	}
	return nil
}
