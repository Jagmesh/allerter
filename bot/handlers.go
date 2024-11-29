package bot

import (
	"fmt"
	"math/rand"
	"strings"
	"tg-alerter/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) HandleUpdate(update tgbotapi.Update) {
	if strings.HasPrefix(update.Message.Text, "@all") || strings.HasPrefix(update.Message.Text, "/all") {
		b.handleAllCommand(update)
	}
}

func (b *Bot) handleAllCommand(update tgbotapi.Update) {
	log := logger.GetLogger()

	admins, err := b.API.GetChatAdministrators(tgbotapi.ChatAdministratorsConfig{ChatConfig: tgbotapi.ChatConfig{ChatID: update.Message.Chat.ID}})
	if err != nil {
		log.Printf("Error getting list of admins: %v", err)
		return
	}
	userName := update.Message.From.UserName
	log.Info("User that trigered command: ", userName)

	var mentions []string
	for _, admin := range admins {
		if admin.User.UserName != "" && admin.User.UserName != b.API.Self.UserName && admin.User.UserName != userName {
			mentions = append(mentions, fmt.Sprintf("@%s", admin.User.UserName))
		}
	}

	mentionText := fmt.Sprintf("@%s зовет ебашить! \n\n%s %s",
		update.Message.From.UserName,
		strings.Join(mentions, " "),
		ListOfMessages[rand.Intn(len(ListOfMessages))])
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, mentionText)
	msg.ParseMode = "Markdown"

	if _, err := b.API.Send(msg); err != nil {
		log.Printf("Error sending message: %v", err)
	}
}
